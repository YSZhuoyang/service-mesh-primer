# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.13

WORKDIR /go/server

# Copy the local package files to the container's workspace.
ADD . .

# install protobuf compiler
RUN apt update && \
    apt -y install apt-utils unzip
RUN PROTOC_ZIP=protoc-3.10.0-linux-x86_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.10.0/protoc-3.10.0-linux-x86_64.zip && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

# Install envoy
RUN apt update && apt -y install \
  software-properties-common \
  curl \
  ca-certificates \
  apt-transport-https \
  gnupg2
RUN curl -sL 'https://getenvoy.io/gpg' | apt-key add -
RUN add-apt-repository "deb [arch=amd64] https://dl.bintray.com/tetrate/getenvoy-deb \
  $(lsb_release -cs) stable"
RUN apt update && apt install -y getenvoy-envoy

# Clone googleapis
RUN cd ../ && git clone https://github.com/googleapis/googleapis

# Install go dependencies and build
RUN make get
RUN make all

# Run the exec file when image is deployed.
RUN chmod u+x ./start_service.sh
ENTRYPOINT ./start_service.sh
