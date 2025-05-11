#!/bin/bash

HOST=127.0.0.1
if [ "$(uname)" = "Linux" ]; then # Debian container
    HOST=host.docker.internal
fi

# HTTP1.1
curl -X POST -d '{"msg": "Hello go service via HTTP/1.1"}' -H "Content-Type:application/json" http://${HOST}:80/greet-go/hello
curl -X POST -d '{"msg": "Hello dotnet service via HTTP/1.1"}' -H "Content-Type:application/json" http://${HOST}:80/greet-dotnet/hello

# HTTP2
curl --http2-prior-knowledge -X POST -d '{"msg": "Hello go service via HTTP/2"}' -H "Content-Type:application/json" http://${HOST}:80/greet-go/hello
curl --http2-prior-knowledge -X POST -d '{"msg": "Hello dotnet service via HTTP/2"}' -H "Content-Type:application/json" http://${HOST}:80/greet-dotnet/hello

# gRPC
grpcurl -plaintext -d '{"msg": "Hello go service gRPC"}' ${HOST}:80 go_service.Greeter/SayHello
grpcurl -plaintext -d '{"msg": "Hello dotnet service gRPC"}' ${HOST}:80 dotnet_service.Greeter/SayHello
