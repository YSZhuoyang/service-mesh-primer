FROM envoyproxy/envoy:v1.12.2

RUN apt update && apt -q install -y curl make

# install protobuf compiler
RUN apt update && \
    apt -y install apt-utils unzip
RUN PROTOC_ZIP=protoc-3.10.0-linux-x86_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.10.0/protoc-3.10.0-linux-x86_64.zip && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

ADD . .

RUN make build

# RUN chmod u+x ./start_service.sh
# ENTRYPOINT ./start_service.sh
ENTRYPOINT [ "envoy", "-c", "./gateway_envoy.yaml" ]
