#!/bin/sh

# Launch asp.net core service with envoy proxy overwriting the cluster name in the bootstrap config file
envoy -c ./gateway-envoy.yaml
