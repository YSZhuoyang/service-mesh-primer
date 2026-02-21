# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
ARG VARIANT="1-bookworm"
FROM mcr.microsoft.com/devcontainers/go:${VARIANT}

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Configure apt, install packages and tools
RUN apt update \
    && apt -y install --no-install-recommends apt-utils unzip dialog 2>&1 \
    && apt -y install git iproute2 procps lsb-release make software-properties-common gnupg

# Install grpcurl
RUN wget https://github.com/fullstorydev/grpcurl/releases/download/v1.9.3/grpcurl_1.9.3_linux_arm64.tar.gz \
    && tar -xvf grpcurl_1.9.3_linux_arm64.tar.gz -C /usr/local/bin \
    && rm grpcurl_1.9.3_linux_arm64.tar.gz

# Install istioctl
ARG ISTIO_VERSION=1.29.0
RUN cd /usr && curl -L https://istio.io/downloadIstio | ISTIO_VERSION=${ISTIO_VERSION} sh -
ENV PATH=$PATH:/usr/istio-${ISTIO_VERSION}/bin

# Copy the kube config sync script
COPY copy-kube-config.sh /usr/local/share/copy-kube-config.sh
RUN chmod +x /usr/local/share/copy-kube-config.sh \
    && echo "source /usr/local/share/copy-kube-config.sh" >> /home/vscode/.bashrc

# Upgrade and clean up
RUN apt upgrade -y && apt autoremove -y --purge && apt autoclean

ENV OS=DEBIAN
# Switch back to dialog for any ad-hoc use of apt
ENV DEBIAN_FRONTEND=dialog
