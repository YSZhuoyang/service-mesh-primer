#!/bin/sh

# Launch asp.net core service with envoy proxy overwriting the cluster name in the bootstrap config file
./server_exec &
envoy -c ./sidecar-envoy.yaml
