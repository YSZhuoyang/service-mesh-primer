# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

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

# install protobuf compiler
RUN PROTOC_VER=21.9 && \
    PROTOC_ZIP=protoc-${PROTOC_VER}-linux-aarch_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/${PROTOC_ZIP} && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

# Install kubectl
RUN curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | tee /etc/apt/sources.list.d/kubernetes.list
RUN apt update && apt install -y kubectl

COPY copy-kube-config.sh /usr/local/share/
RUN echo "source /usr/local/share/copy-kube-config.sh" | tee -a /home/vscode/.bashrc >> /home/vscode/.zshrc

# Install istioctl
ARG ISTIO_VERSION=1.15.3
RUN cd /usr && curl -L https://istio.io/downloadIstio | ISTIO_VERSION=${ISTIO_VERSION} sh -
ENV PATH=$PATH:/usr/istio-${ISTIO_VERSION}/bin

# Upgrade and clean up
RUN apt upgrade -y && apt autoremove -y --purge && apt autoclean

ENV OS=DEBIAN
# Switch back to dialog for any ad-hoc use of apt
ENV DEBIAN_FRONTEND=dialog
