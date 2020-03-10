FROM golang:1.14
WORKDIR /go/server

# install protobuf compiler
RUN apt update && \
    apt -y install apt-utils unzip
RUN PROTOC_ZIP=protoc-3.11.4-linux-x86_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

# Install envoy
RUN apt update && apt -y install \
  apt-utils \
  software-properties-common \
  curl \
  ca-certificates \
  apt-transport-https \
  gnupg2 make
RUN curl -sL 'https://getenvoy.io/gpg' | apt-key add -
RUN add-apt-repository "deb [arch=amd64] https://dl.bintray.com/tetrate/getenvoy-deb \
  $(lsb_release -cs) stable"
RUN apt update && apt install -y getenvoy-envoy

# Build src
COPY ./ ./
RUN make all

# Start service with envoy sidecar
RUN chmod u+x ./start_service.sh
ENTRYPOINT ./start_service.sh
