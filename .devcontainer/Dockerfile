# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
ARG VARIANT="1"
FROM mcr.microsoft.com/devcontainers/go:${VARIANT}

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Configure apt, install packages and tools
RUN apt update \
    && apt -y install --no-install-recommends apt-utils unzip dialog 2>&1 \
    && apt -y install git iproute2 procps lsb-release make software-properties-common gnupg

# Install kubectl
# RUN apt -y install --no-install-recommends \
#         # apt-transport-https may be a dummy package; if so, you can skip that package
#         apt-transport-https \
#         ca-certificates \
#         curl \
#     && curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.28/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg

# This overwrites any existing configuration in /etc/apt/sources.list.d/kubernetes.list
# RUN echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.28/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
# RUN apt update && apt install -y kubectl

# Install Nodejs TLS
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - \
    && apt install -y nodejs

# Install istioctl
ARG ISTIO_VERSION=1.24.1
RUN cd /usr && curl -L https://istio.io/downloadIstio | ISTIO_VERSION=${ISTIO_VERSION} sh -
ENV PATH=$PATH:/usr/istio-${ISTIO_VERSION}/bin

# Upgrade and clean up
RUN apt upgrade -y && apt autoremove -y --purge && apt autoclean

ENV OS=DEBIAN
# Switch back to dialog for any ad-hoc use of apt
ENV DEBIAN_FRONTEND=dialog
