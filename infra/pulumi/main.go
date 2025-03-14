package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

const (
	VNET_RANGE           = "10.0.0.0/16"
	DB_SUBNET_RANGE      = "10.0.1.0/24"
	BACKEND_SUBNET_RANGE = "10.0.2.0/24"
	VM_SUBNET_RANGE      = "10.0.3.0/24"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		resourceGroup, err := resources.NewResourceGroup(ctx, "foxnhound-rg", nil)
		if err != nil {
			return err
		}

		vnet, err := network.NewVirtualNetwork(ctx, "foxnhound-vnet", &network.VirtualNetworkArgs{
			ResourceGroupName: resourceGroup.Name,
			AddressSpace: &network.AddressSpaceArgs{
				AddressPrefixes: pulumi.StringArray{
					pulumi.String(VNET_RANGE),
				},
			},
		})
		if err != nil {
			return err
		}

		// Create a SQL Server
		// _, err = createSqlServer(ctx, resourceGroup, subnet)
		// if err != nil {
		// 	return err
		// }

		// Create MySQL Server
		serverReturn := createMySqlServer(MySqlServerArgs{
			ctx:           ctx,
			resourceGroup: resourceGroup,
			vnet:          vnet})
		if serverReturn.err != nil || serverReturn.dbServer == nil || serverReturn.db == nil {
			return serverReturn.err
		}

		// Create a Web App running a container on Linux
		backendReturn := createBackend(BackendArgs{
			ctx:           ctx,
			resourceGroup: resourceGroup,
			vnet:          vnet,
			dbServerFQDN:  serverReturn.dbServer.FullyQualifiedDomainName,
			dbName:        serverReturn.db.Name,
		})
		if backendReturn.err != nil {
			return backendReturn.err
		}

		frontendReturn := createFrontend(FrontendArgs{
			ctx:             ctx,
			resourceGroup:   resourceGroup,
			appServicePlan:  backendReturn.appServicePlan,
			defaultHostname: backendReturn.webapp.DefaultHostName,
		})
		if frontendReturn.err != nil {
			return frontendReturn.err
		}

		vmReturn := createVm(VmArgs{
			ctx:           ctx,
			resourceGroup: resourceGroup,
			vnet:          vnet,
		})
		if vmReturn.err != nil {
			return vmReturn.err
		}

		return nil
	})

}

type VmArgs struct {
	ctx           *pulumi.Context
	resourceGroup *resources.ResourceGroup
	vnet          *network.VirtualNetwork
}

type VmReturn struct {
	vm  *compute.VirtualMachine
	err error
}

func createVm(args VmArgs) VmReturn {
	config := config.New(args.ctx, "")

	vmSubnet, err := network.NewSubnet(args.ctx, "foxnhound-vm-subnet", &network.SubnetArgs{
		Name:               pulumi.String("internal"),
		ResourceGroupName:  args.resourceGroup.Name,
		VirtualNetworkName: args.vnet.Name,
		AddressPrefixes: pulumi.StringArray{
			pulumi.String(VM_SUBNET_RANGE),
		},
	})
	if err != nil {
		return VmReturn{err: err}
	}

	vmPublicIp, err := network.NewPublicIPAddress(args.ctx, "foxnhound-vm-public-ip", &network.PublicIPAddressArgs{
		ResourceGroupName: args.resourceGroup.Name,
	})
	if err != nil {
		return VmReturn{err: err}
	}

	vmNetworkInterface, err := network.NewNetworkInterface(args.ctx, "foxnhound-vm-network-interface", &network.NetworkInterfaceArgs{
		IpConfigurations: network.NetworkInterfaceIPConfigurationArray{
			&network.NetworkInterfaceIPConfigurationArgs{
				Name: pulumi.String("ipconfig1"),
				PublicIPAddress: &network.PublicIPAddressTypeArgs{
					Id: vmPublicIp.ID(),
				},
				Subnet: &network.SubnetTypeArgs{
					Id: vmSubnet.ID(),
				},
			},
		},
		ResourceGroupName: args.resourceGroup.Name,
	})
	if err != nil {
		return VmReturn{err: err}
	}

	adminLogin := config.Require("vm-adminLogin")
	adminSecret := config.Require("vm-adminSecret")
	vm, err := compute.NewVirtualMachine(args.ctx, "foxnhound-vm", &compute.VirtualMachineArgs{
		ResourceGroupName: args.resourceGroup.Name,
		NetworkProfile: compute.NetworkProfileArgs{
			NetworkInterfaces: &compute.NetworkInterfaceReferenceArray{
				&compute.NetworkInterfaceReferenceArgs{
					Id: vmNetworkInterface.ID(),
				},
			},
		},
		HardwareProfile: &compute.HardwareProfileArgs{
			VmSize: pulumi.String("Standard_DS1_v2"),
		},
		OsProfile: &compute.OSProfileArgs{
			ComputerName:  pulumi.String("foxnhound-vm"),
			AdminUsername: pulumi.String(adminLogin),
			AdminPassword: pulumi.String(adminSecret),
		},
		StorageProfile: &compute.StorageProfileArgs{
			OsDisk: compute.OSDiskArgs{
				CreateOption: pulumi.String("FromImage"),
				ManagedDisk: &compute.ManagedDiskParametersArgs{
					StorageAccountType: pulumi.String("Standard_LRS"),
				},
			},
			ImageReference: &compute.ImageReferenceArgs{
				Publisher: pulumi.String("Canonical"),
				Offer:     pulumi.String("UbuntuServer"),
				Sku:       pulumi.String("18.04-LTS"),
				Version:   pulumi.String("latest"),
			},
		},
	})
	if err != nil {
		return VmReturn{err: err}
	}

	return VmReturn{vm: vm}
}

type BackendArgs struct {
	ctx           *pulumi.Context
	resourceGroup *resources.ResourceGroup
	vnet          *network.VirtualNetwork
	dbServerFQDN  pulumi.StringOutput
	dbName        pulumi.StringOutput
}

type BackendReturn struct {
	webapp         *web.WebApp
	appServicePlan *web.AppServicePlan
	err            error
}

func createBackend(args BackendArgs) BackendReturn {
	config := config.New(args.ctx, "")

	subnet, err := network.NewSubnet(args.ctx, "foxnhound-sn-backend", &network.SubnetArgs{
		Delegations: network.DelegationArray{
			&network.DelegationArgs{
				Name:        pulumi.String("backendDelegation"),
				ServiceName: pulumi.String("Microsoft.Web/serverFarms"),
			},
		},
		ResourceGroupName:  args.resourceGroup.Name,
		VirtualNetworkName: args.vnet.Name,
		AddressPrefix:      pulumi.String(BACKEND_SUBNET_RANGE),
	})
	if err != nil {
		return BackendReturn{err: err}
	}

	// Create an App Service Plan if wen need to
	appServicePlan, err := web.NewAppServicePlan(args.ctx, "appServicePlan", &web.AppServicePlanArgs{
		ResourceGroupName: args.resourceGroup.Name,
		Location:          args.resourceGroup.Location,
		Sku: &web.SkuDescriptionArgs{
			Name:     pulumi.String("B1"),
			Tier:     pulumi.String("Basic"),
			Capacity: pulumi.Int(1),
		},
		Reserved: pulumi.Bool(true), // Reserved indicates Linux
	})
	if err != nil {
		return BackendReturn{err: err}
	}

	containerRegistryLogin := config.Require("container-registry-login")
	containerRegistryPassword := config.Require("container-registry-password")
	containerRegistryUrl := config.Require("container-registry-url")
	containerRegistryBasePath := config.Require("container-registry-base-path")
	backendImageTag := config.Require("backend-image-tag")
	webApp, err := web.NewWebApp(args.ctx, "foxnhound-backend", &web.WebAppArgs{
		ResourceGroupName: args.resourceGroup.Name,
		Location:          args.resourceGroup.Location,
		ServerFarmId:      appServicePlan.ID(),
		SiteConfig: &web.SiteConfigArgs{
			AlwaysOn:       pulumi.Bool(true),
			LinuxFxVersion: pulumi.String(containerRegistryBasePath + ":" + backendImageTag),
			AppSettings: web.NameValuePairArray{
				&web.NameValuePairArgs{
					Name:  pulumi.String("DOCKER_REGISTRY_SERVER_URL"),
					Value: pulumi.String(containerRegistryUrl),
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("DOCKER_REGISTRY_SERVER_USERNAME"),
					Value: pulumi.String(containerRegistryLogin),
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("DOCKER_REGISTRY_SERVER_PASSWORD"),
					Value: pulumi.String(containerRegistryPassword),
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("DB_USER"),
					Value: pulumi.String(config.Require("mysql-adminLogin")),
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("DB_PASSWORD"),
					Value: pulumi.String(config.Require("mysql-adminSecret")),
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("DB_HOST"),
					Value: args.dbServerFQDN,
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("DB_PORT"),
					Value: pulumi.String("3306"),
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("DB_NAME"),
					Value: args.dbName,
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("DB_TLS"),
					Value: pulumi.String("CUSTOM"),
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("CA_CERT_PATH"),
					Value: pulumi.String("certs/DigiCertGlobalRootCA.crt.pem"),
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("STAGE"),
					Value: pulumi.String("dev"),
				},
			},
		},
		VirtualNetworkSubnetId: subnet.ID(),
	})

	return BackendReturn{webapp: webApp, appServicePlan: appServicePlan}
}

type FrontendArgs struct {
	ctx             *pulumi.Context
	resourceGroup   *resources.ResourceGroup
	appServicePlan  *web.AppServicePlan
	defaultHostname pulumi.StringOutput
}

type FrontendReturn struct {
	webapp *web.WebApp
	err    error
}

func createFrontend(args FrontendArgs) FrontendReturn {
	config := config.New(args.ctx, "")
	containerRegistryLogin := config.Require("container-registry-login")
	containerRegistryPassword := config.Require("container-registry-password")
	containerRegistryUrl := config.Require("container-registry-url")
	containerRegistryBasePath := config.Require("container-registry-base-path")
	webappImageTag := config.Require("webapp-image-tag")
	backendUrl := pulumi.Sprintf("https://%s", args.defaultHostname)

	webApp, err := web.NewWebApp(args.ctx, "foxnhound-webapp", &web.WebAppArgs{
		ResourceGroupName: args.resourceGroup.Name,
		Location:          args.resourceGroup.Location,
		ServerFarmId:      args.appServicePlan.ID(),
		SiteConfig: &web.SiteConfigArgs{
			AlwaysOn:       pulumi.Bool(true),
			LinuxFxVersion: pulumi.String(containerRegistryBasePath + ":" + webappImageTag),
			AppSettings: web.NameValuePairArray{
				&web.NameValuePairArgs{
					Name:  pulumi.String("DOCKER_REGISTRY_SERVER_URL"),
					Value: pulumi.String(containerRegistryUrl),
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("DOCKER_REGISTRY_SERVER_USERNAME"),
					Value: pulumi.String(containerRegistryLogin),
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("DOCKER_REGISTRY_SERVER_PASSWORD"),
					Value: pulumi.String(containerRegistryPassword),
				},
				&web.NameValuePairArgs{
					Name: pulumi.String("BACKEND_URL"),
					Value: backendUrl,
				},
				&web.NameValuePairArgs{
					Name:  pulumi.String("STAGE"),
					Value: pulumi.String("dev"),
				},
			},
		},
	})
	if err != nil {
		return FrontendReturn{err: err}
	}

	return FrontendReturn{webapp: webApp}
}

type MySqlServerArgs struct {
	ctx           *pulumi.Context
	resourceGroup *resources.ResourceGroup
	vnet          *network.VirtualNetwork
}

type MySqlServerReturn struct {
	dbServer *dbformysql.Server
	db       *dbformysql.Database
	err      error
}

func createMySqlServer(args MySqlServerArgs) MySqlServerReturn {
	conf := config.New(args.ctx, "")

	// Cretae a new delegated subnet
	subnet, err := network.NewSubnet(args.ctx, "foxnhound-sn-db", &network.SubnetArgs{
		Delegations: network.DelegationArray{
			&network.DelegationArgs{
				ServiceName: pulumi.String("Microsoft.DBforMySQL/flexibleServers"),
				Name:        pulumi.String("mysqlDelegation"),
			},
		},
		ResourceGroupName:  args.resourceGroup.Name,
		VirtualNetworkName: args.vnet.Name,
		AddressPrefix:      pulumi.String(DB_SUBNET_RANGE),
	})
	if err != nil {
		return MySqlServerReturn{err: err}
	}

	dnszone, err := network.NewPrivateZone(args.ctx, "foxnhound-private-dns-zone", &network.PrivateZoneArgs{
		ResourceGroupName: args.resourceGroup.Name,
		PrivateZoneName:   pulumi.String("foxnhound.mysql.database.azure.com"),
		Location:          pulumi.String("Global"),
	})
	if err != nil {
		return MySqlServerReturn{err: err}
	}

	networklink, err := network.NewVirtualNetworkLink(args.ctx, "foxnhound-db-nwlink", &network.VirtualNetworkLinkArgs{
		ResourceGroupName:   args.resourceGroup.Name,
		PrivateZoneName:     dnszone.Name,
		Location:            pulumi.String("Global"),
		RegistrationEnabled: pulumi.Bool(false),
		VirtualNetwork: &network.SubResourceArgs{
			Id: args.vnet.ID(),
		},
	})
	if err != nil {
		return MySqlServerReturn{err: err}
	}

	adminLogin := conf.Require("mysql-adminLogin")
	adminSecret := conf.Require("mysql-adminSecret")
	dbserver, err := dbformysql.NewServer(args.ctx, "foxnhound-mysql-server", &dbformysql.ServerArgs{
		AdministratorLogin:         pulumi.String(adminLogin),
		AdministratorLoginPassword: pulumi.String(adminSecret),
		ResourceGroupName:          args.resourceGroup.Name,
		Sku: &dbformysql.SkuArgs{
			Name: pulumi.String("Standard_B1ms"),
			Tier: pulumi.String(dbformysql.SkuTierBurstable),
		},
		Version: pulumi.String(dbformysql.ServerVersion_8_0_21),
		Network: &dbformysql.NetworkArgs{
			DelegatedSubnetResourceId: subnet.ID(),
			PrivateDnsZoneResourceId:  dnszone.ID(),
		},
	},
		pulumi.DependsOn([]pulumi.Resource{subnet, dnszone, networklink}))
	if err != nil {
		return MySqlServerReturn{err: err}
	}

	// Create a private DNS CNAME record for the MySQL server
	// TODO: Would be sexy to have a fixed dns for the db but requires a self-signed cert, meh
	// _, err = network.NewPrivateRecordSet(args.ctx, "foxnhound-mysql-dns-record", &network.PrivateRecordSetArgs{
	// 	CnameRecord: &network.CnameRecordArgs{
	// 		Cname: dbserver.FullyQualifiedDomainName,
	// 	},
	// 	PrivateZoneName:       dnszone.Name,
	// 	RecordType:            pulumi.String("CNAME"),
	// 	RelativeRecordSetName: pulumi.String("mysql"),
	// 	ResourceGroupName:     args.resourceGroup.Name,
	// 	Ttl:                   pulumi.Float64(300),
	// }, pulumi.DependsOn([]pulumi.Resource{dbserver, dnszone}))
	// if err != nil {
	// 	return MySqlServerReturn{err: err}
	// }

	db, err := dbformysql.NewDatabase(args.ctx, "foxnhound-db", &dbformysql.DatabaseArgs{
		Charset:           pulumi.String("utf8"),
		Collation:         pulumi.String("utf8_general_ci"),
		ServerName:        dbserver.Name,
		ResourceGroupName: args.resourceGroup.Name,
	})
	if err != nil {
		return MySqlServerReturn{err: err}
	}

	return MySqlServerReturn{dbServer: dbserver, db: db}

}

type SqlServerArgs struct {
	ctx           *pulumi.Context
	resourceGroup *resources.ResourceGroup
	vnet          *network.VirtualNetwork
}

type SqlServerReturn struct {
	dbserver *sql.Server
	err      error
}

func createSqlServer(args SqlServerArgs) SqlServerReturn {
	config := config.New(args.ctx, "")

	subnet, err := network.NewSubnet(args.ctx, "foxnhound-subnet", &network.SubnetArgs{
		ResourceGroupName:  args.resourceGroup.Name,
		VirtualNetworkName: args.vnet.Name,
		AddressPrefix:      pulumi.String(DB_SUBNET_RANGE),
	})
	if err != nil {
		return SqlServerReturn{err: err}
	}

	azureUserEMail := config.Require("azureUserEMail")
	azureUserSid := config.Require("azureUserSid")
	azureUserTenantId := config.Require("azureUserTenantId")
	dbserver, err := sql.NewServer(args.ctx, "foxnhound-db-server", &sql.ServerArgs{
		Administrators: &sql.ServerExternalAdministratorArgs{
			AzureADOnlyAuthentication: pulumi.Bool(true),
			Login:                     pulumi.String(azureUserEMail),
			PrincipalType:             pulumi.String(sql.PrincipalTypeUser),
			Sid:                       pulumi.String(azureUserSid),
			TenantId:                  pulumi.String(azureUserTenantId),
		},
		PublicNetworkAccess:           pulumi.String(sql.ServerNetworkAccessFlagDisabled),
		ResourceGroupName:             args.resourceGroup.Name,
		RestrictOutboundNetworkAccess: pulumi.String(sql.ServerNetworkAccessFlagDisabled),
	})
	if err != nil {
		return SqlServerReturn{err: err}
	}

	_, err = sql.NewDatabase(args.ctx, "foxnhound-db", &sql.DatabaseArgs{
		ResourceGroupName: args.resourceGroup.Name,
		ServerName:        dbserver.Name,
		Sku: &sql.SkuArgs{
			Name: pulumi.String("Basic"),
			Tier: pulumi.String("Basic"),
		},
	})
	if err != nil {
		return SqlServerReturn{err: err}
	}

	// Create a Private Endpoint for the Azure SQL Database
	privateEndpoint, err := network.NewPrivateEndpoint(args.ctx, "foxnhound-private-endpoint", &network.PrivateEndpointArgs{
		ResourceGroupName: args.resourceGroup.Name,
		Subnet: &network.SubnetTypeArgs{
			Id: subnet.ID(),
		},
		PrivateLinkServiceConnections: network.PrivateLinkServiceConnectionArray{
			&network.PrivateLinkServiceConnectionArgs{
				Name:                 pulumi.String("sqlPrivateLink"),
				PrivateLinkServiceId: dbserver.ID(),
				GroupIds: pulumi.StringArray{
					pulumi.String("sqlServer"),
				},
				RequestMessage: pulumi.String("Please approve my connection"),
			},
		},
	})
	if err != nil {
		return SqlServerReturn{err: err}
	}

	// Create a Private DNS Zone for the SQL Server
	privateDnsZone, err := network.NewPrivateZone(args.ctx, "foxnhound-private-dns-zone", &network.PrivateZoneArgs{
		ResourceGroupName: args.resourceGroup.Name,
		PrivateZoneName:   pulumi.String("foxnhound.privatelink.database.windows.net"),
		Location:          pulumi.String("Global"),
	})
	if err != nil {
		return SqlServerReturn{err: err}
	}

	// Create a DNS Zone Group for the Private Endpoint
	_, err = network.NewPrivateDnsZoneGroup(args.ctx, "foxnhound-dns-zone-group", &network.PrivateDnsZoneGroupArgs{
		ResourceGroupName:   args.resourceGroup.Name,
		PrivateEndpointName: privateEndpoint.Name,
		PrivateDnsZoneConfigs: network.PrivateDnsZoneConfigArray{
			&network.PrivateDnsZoneConfigArgs{
				Name:             pulumi.String("sqlDnsZoneConfig"),
				PrivateDnsZoneId: privateDnsZone.ID(),
			},
		},
	})
	if err != nil {
		return SqlServerReturn{err: err}
	}

	return SqlServerReturn{dbserver: dbserver}
}
