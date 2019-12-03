#!/bin/sh

# Launch asp.net core service with envoy proxy overwriting the cluster name in the bootstrap config file
dotnet server.dll &
envoy -c /etc/service-proxy-envoy.yaml --service-cluster service${SERVICE_NAME}
