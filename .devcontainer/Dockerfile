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

# Install Register Microsoft key and feed for .NET
RUN wget https://packages.microsoft.com/config/debian/10/packages-microsoft-prod.deb -O packages-microsoft-prod.deb \
    && dpkg -i packages-microsoft-prod.deb

# Install .NET SDK
RUN apt update && apt -y install dotnet-sdk-5.0

# Install Docker CLI & K8s
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
RUN echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
RUN apt update && apt install -y docker-ce-cli

# Install kubectl
RUN curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
RUN apt update && apt install -y kubectl

# install protobuf compiler
RUN PROTOC_VER=3.15.8 && \
    PROTOC_ZIP=protoc-${PROTOC_VER}-linux-x86_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/${PROTOC_ZIP} && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

# Install istioctl
ARG ISTIO_VERSION="1.9.3"
RUN curl -L https://istio.io/downloadIstio | ISTIO_VERSION=${ISTIO_VERSION} sh -
ENV PATH="${PWD}/istio-${ISTIO_VERSION}/bin:${PATH}"

# Upgrade and clean up
RUN apt upgrade -y && apt autoremove -y --purge && apt autoclean

# Switch back to dialog for any ad-hoc use of apt
ENV DEBIAN_FRONTEND=dialog
