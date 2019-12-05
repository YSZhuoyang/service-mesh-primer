#!/bin/bash

# echo "-------- Starting GRPC server -----------"
# ./bin/personserver &
# echo "-------- Starting Envoy -----------"
# envoy -c ./conf/envoy.yaml &

# sleep 5

echo "-------- Running tests -----------"
curl -X POST -d '{"name": "Sean", "age": 21}' -H "Content-Type: application/json" http://0.0.0.0:7778/create
curl -X POST -d '{"name": "Jeff", "age": 42}' -H "Content-Type: application/json" http://0.0.0.0:7778/create
curl -X POST -d '{"name": "Jeff", "age": 42}' -H "Content-Type: application/json" http://0.0.0.0:7778/create
curl -H "Content-Type: application/json" http://0.0.0.0:7778/lookup?name=Sean\&age=21
curl -H "Content-Type: application/json" http://0.0.0.0:7778/lookup?name=Jeff\&age=42
curl -H "Content-Type: application/json" http://0.0.0.0:7778/lookup?name=Neil\&age=42
