#!/bin/sh
dotnet server.dll &
envoy -c /etc/service-proxy-envoy.yaml --service-cluster service${SERVICE_NAME}
# Overwrite the cluster name in the bootstrap config file
