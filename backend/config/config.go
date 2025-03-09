package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbTLSSetting string

const (
	DbTlsDisable DbTLSSetting = "DISABLE"
	DbTlsCustom  DbTLSSetting = "CUSTOM"
)

type Stage string

const (
	StageLocal Stage = "local"
	StageDev   Stage = "dev"
	StageProd  Stage = "prod"
)

type Config struct {
	stage        Stage
	dbUser       string
	dbPassword   string
	dbHost       string
	dbPort       string
	dbName       string
	dbTlsSetting DbTLSSetting
	caCertPath   string
}

var config Config
var DB *gorm.DB

func init() {
	stage := os.Getenv("STAGE")
	if stage != string(StageDev) && stage != string(StageProd) {
		log.Printf("Not in a remote environment, loading .env file")
		loadEnv()
	}
	config = buidConfig()
}

func buidConfig() Config {
	config := getDefaultConfig()
	// overwrite the default values with the environment variables
	if os.Getenv("DB_USER") != "" {
		log.Printf("Found env variable DB_USER: %s", os.Getenv("DB_USER"))
		config.dbUser = os.Getenv("DB_USER")
	}
	if os.Getenv("DB_PASSWORD") != "" {
		log.Printf("Found env variable DB_PASSWORD: -----")
		config.dbPassword = os.Getenv("DB_PASSWORD")
	}
	if os.Getenv("DB_HOST") != "" {
		log.Printf("Found env variable DB_HOST: %s", os.Getenv("DB_HOST"))
		config.dbHost = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_PORT") != "" {
		log.Printf("Found env variable DB_PORT: %s", os.Getenv("DB_PORT"))
		config.dbPort = os.Getenv("DB_PORT")
	}
	if os.Getenv("DB_NAME") != "" {
		log.Printf("Found env variable DB_NAME: %s", os.Getenv("DB_NAME"))
		config.dbName = os.Getenv("DB_NAME")
	}
	if os.Getenv("DB_TLS") != "" {
		log.Printf("Found env variable DB_TLS: %s", os.Getenv("DB_TLS"))
		config.dbTlsSetting = DbTLSSetting(os.Getenv("DB_TLS"))
	}
	if os.Getenv("CA_CERT_PATH") != "" {
		log.Printf("Found env variable CA_CERT_PATH: %s", os.Getenv("CA_CERT_PATH"))
		config.caCertPath = os.Getenv("CA_CERT_PATH")
	}
	if os.Getenv("STAGE") != "" {
		log.Printf("Found env variable STAGE: %s", os.Getenv("STAGE"))
		config.stage = Stage(os.Getenv("STAGE"))
	}
	return config
}

func getDefaultConfig() Config {
	return Config{
		stage:        StageProd, // default to prod as this should be the safe case
		dbUser:       "devuser",
		dbPassword:   "devpassword",
		dbHost:       "127.0.0.1",
		dbPort:       "3306",
		dbName:       "fox_and_hound",
		dbTlsSetting: DbTlsDisable,
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
}

func addCustomCert() {
	// Load the CA certificate
	caCert, err := os.ReadFile(config.caCertPath)
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Register a custom TLS config
	err = mysqlDriver.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: caCertPool,
	})
	if err != nil {
		log.Fatalf("Failed to register custom TLS config: %v", err)
		return
	}
}

func ConnectDatabase() *gorm.DB {

	if config.dbTlsSetting == DbTlsCustom {
		addCustomCert()
	}

	// DNS resolution debug code
	addrs, err := net.LookupHost(config.dbHost)
	if err != nil {
		log.Fatalf("Failed to resolve hostname %s: %v", config.dbHost, err)
	}
	log.Printf("Resolved hostname %s to addresses: %v", config.dbHost, addrs)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.dbUser,
		config.dbPassword,
		config.dbHost,
		config.dbPort,
		config.dbName,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}))

	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Database connected successfully")

	DB = db
	return db
}
