#!/bin/bash
# Change "127.0.0.1" to "host.docker.internal" to run this script from vscode dev container

# HTTP1.1
curl -X POST -d '{"msg": "Hello go service"}' -H "Content-Type:application/json" http://127.0.0.1:80/greet/hello
curl -X POST -d '{"msg": "Hello dotnet service"}' -H "Content-Type:application/json" http://127.0.0.1:80/greet2/hello

# HTTP2
curl --http2-prior-knowledge -X POST -d '{"msg": "Hello go service"}' -H "Content-Type:application/json" http://127.0.0.1:80/greet/hello
curl --http2-prior-knowledge -X POST -d '{"msg": "Hello dotnet service"}' -H "Content-Type:application/json" http://127.0.0.1:80/greet2/hello
