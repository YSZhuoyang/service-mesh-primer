#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

FROM mcr.microsoft.com/vscode/devcontainers/go:1

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Configure apt, install packages and tools
RUN apt update \
    && apt -y install --no-install-recommends apt-utils unzip dialog 2>&1 \
    && apt -y install git iproute2 procps lsb-release make software-properties-common \
        apt-transport-https ca-certificates curl gnupg

# Install Nodejs TLS
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - \
    && apt install -y nodejs

# Install Docker CLI
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
    $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
RUN apt update && apt install -y docker-ce-cli

# Install kubectl
RUN curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | tee /etc/apt/sources.list.d/kubernetes.list
RUN apt update && apt install -y kubectl

# install protobuf compiler
RUN PROTOC_VER=21.1 && \
    PROTOC_ZIP=protoc-${PROTOC_VER}-linux-aarch_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/${PROTOC_ZIP} && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

# Install .NET SDK (available in arm64), see https://docs.microsoft.com/en-gb/dotnet/core/install/linux-scripted-manual#scripted-install
RUN wget https://dot.net/v1/dotnet-install.sh \
    && chmod +x dotnet-install.sh \
    && bash ./dotnet-install.sh -c 6.0
RUN ln -s ~/.dotnet /usr/local/bin/dotnet
ENV PATH=$PATH:/usr/local/bin/dotnet

# Install istioctl
ARG ISTIO_VERSION=1.14.1
RUN cd /usr && curl -L https://istio.io/downloadIstio | ISTIO_VERSION=${ISTIO_VERSION} sh -
ENV PATH=$PATH:/usr/istio-${ISTIO_VERSION}/bin

# Upgrade and clean up
RUN apt upgrade -y && apt autoremove -y --purge && apt autoclean

ENV OS=DEBIAN
# Switch back to dialog for any ad-hoc use of apt
ENV DEBIAN_FRONTEND=dialog
