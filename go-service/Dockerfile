FROM golang:1.15
WORKDIR /go/server

# install protobuf compiler
RUN apt update && \
    apt -y install apt-utils unzip
RUN PROTOC_VER=3.14.0 && \
    PROTOC_ZIP=protoc-${PROTOC_VER}-linux-x86_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/${PROTOC_ZIP} && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

# Build src
COPY ./ ./
RUN make all

ENTRYPOINT ./server_exec
