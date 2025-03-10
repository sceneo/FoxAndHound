FROM mcr.microsoft.com/devcontainers/base:bullseye

# Install dependencies
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    gnupg \
    software-properties-common \
    openjdk-11-jdk

# Install Go 1.23.5
RUN wget https://go.dev/dl/go1.23.5.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.23.5.linux-amd64.tar.gz \
    && rm go1.23.5.linux-amd64.tar.gz

ENV PATH "$PATH:/usr/local/go/bin"

# Switch to the vscode user
USER vscode
# Install NVM and Node.js for the vscode user
ENV NODE_VERSION=22
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
ENV NVM_DIR=/home/vscode/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/home/vscode/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"
RUN . "$NVM_DIR/nvm.sh" && npm install -g @angular/cli

# Install OpenAPI Generator CLI
RUN . "$NVM_DIR/nvm.sh" && npm install -g @openapitools/openapi-generator-cli

# Switch back to root to install Azure CLI and other tools
USER root

# Install Azure CLI
RUN curl -sL https://aka.ms/InstallAzureCLIDeb | bash

# Installing pulumi
RUN curl -fsSL https://get.pulumi.com | sh 
RUN mkdir /usr/bin/.pulumi && mv /root/.pulumi/* /usr/bin/.pulumi/ 
ENV PATH "$PATH:/usr/bin/.pulumi/bin/"