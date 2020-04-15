#!/bin/bash

# HTTP1.1
curl --insecure -X POST -d '{"Msg": "Hello go service"}' -H "Content-Type:application/json" https://stock-discovery-oyu.space:443/greet/hello
# curl -X POST -d '{"Msg": "Hello dotnet service"}' -H "Content-Type:application/json" http://0.0.0.0:80/greet2/hello

# HTTP2
curl --insecure --http2-prior-knowledge -X POST -d '{"Msg": "Hello go service"}' -H "Content-Type:application/json" https://stock-discovery-oyu.space:443/greet/hello
# curl --http2-prior-knowledge -X POST -d '{"Msg": "Hello dotnet service"}' -H "Content-Type:application/json" http://0.0.0.0:80/greet2/hello
