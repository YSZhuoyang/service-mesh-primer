# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1-bullseye, 1.19-bullseye
ARG VARIANT="1-trixie"
FROM mcr.microsoft.com/devcontainers/go:${VARIANT}

# Avoid warnings by switching to noninteractive
ARG DEBIAN_FRONTEND=noninteractive

# Configure apt, install packages and tools
RUN apt-get update && apt-get install -y --no-install-recommends \
    apt-utils \
    ca-certificates \
    curl \
    gnupg \
    unzip \
    git \
    iproute2 \
    procps \
    lsb-release \
    make \
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Install grpcurl
RUN wget https://github.com/fullstorydev/grpcurl/releases/download/v1.9.3/grpcurl_1.9.3_linux_arm64.tar.gz \
    && tar -xvf grpcurl_1.9.3_linux_arm64.tar.gz -C /usr/local/bin \
    && rm grpcurl_1.9.3_linux_arm64.tar.gz

# Install istioctl
ARG ISTIO_VERSION=1.29.1
RUN cd /usr && curl -L https://istio.io/downloadIstio | ISTIO_VERSION=${ISTIO_VERSION} sh - \
    && chmod -R +rX /usr/istio-${ISTIO_VERSION}
ENV PATH=$PATH:/usr/istio-${ISTIO_VERSION}/bin

# Copy the kube config sync script
COPY copy-kube-config.sh /usr/local/share/copy-kube-config.sh
RUN chmod +x /usr/local/share/copy-kube-config.sh \
    && echo "source /usr/local/share/copy-kube-config.sh" >> /home/vscode/.bashrc

ENV OS=DEBIAN
# Switch back to dialog for any ad-hoc use of apt
ARG DEBIAN_FRONTEND=dialog
