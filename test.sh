#!/bin/bash

HOST=127.0.0.1
if [ "$(uname)" = "Linux" ]; then # Debian container
    HOST=host.docker.internal
fi

# HTTP1.1
curl -X POST -d '{"msg": "Hello go service"}' -H "Content-Type:application/json" http://${HOST}:80/greet-go/hello:83
curl -X POST -d '{"msg": "Hello dotnet service"}' -H "Content-Type:application/json" http://${HOST}:80/greet-dotnet/hello:83

# HTTP2
curl --http2-prior-knowledge -X POST -d '{"msg": "Hello go service"}' -H "Content-Type:application/json" http://${HOST}:80/greet-go/hello:83
curl --http2-prior-knowledge -X POST -d '{"msg": "Hello dotnet service"}' -H "Content-Type:application/json" http://${HOST}:80/greet-dotnet/hello:83
