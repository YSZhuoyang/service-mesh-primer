#!/bin/bash

# Default to localhost as most port-forwards run in the local environment
HOST=127.0.0.1

# Only use host.docker.internal if we are in a container and 127.0.0.1:80 is not responding 
# (meaning the port-forward is likely on the host machine).
if ! curl -s --connect-timeout 1 http://127.0.0.1:8080/ > /dev/null; then
    if getent hosts host.docker.internal > /dev/null; then
        HOST=host.docker.internal
    fi
fi

echo "Using HOST: ${HOST}"

# HTTP1.1
curl -X POST -d '{"msg": "Hello go service via HTTP/1.1"}' -H "Content-Type:application/json" http://${HOST}:8080/greet-go/hello
curl -X POST -d '{"msg": "Hello dotnet service via HTTP/1.1"}' -H "Content-Type:application/json" http://${HOST}:8080/greet-dotnet/hello

# HTTP2
curl --http2-prior-knowledge -X POST -d '{"msg": "Hello go service via HTTP/2"}' -H "Content-Type:application/json" http://${HOST}:8080/greet-go/hello
curl --http2-prior-knowledge -X POST -d '{"msg": "Hello dotnet service via HTTP/2"}' -H "Content-Type:application/json" http://${HOST}:8080/greet-dotnet/hello

# gRPC
grpcurl -plaintext -d '{"msg": "Hello go service gRPC"}' ${HOST}:8080 go_service.Greeter/SayHello
grpcurl -plaintext -d '{"msg": "Hello dotnet service gRPC"}' ${HOST}:8080 dotnet_service.Greeter/SayHello
