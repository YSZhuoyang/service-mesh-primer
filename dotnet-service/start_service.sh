#!/bin/sh

# Launch asp.net core service with envoy proxy overwriting the cluster name in the bootstrap config file
dotnet server.dll &
envoy -c /etc/sidecar-envoy.yaml
