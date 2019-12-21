#!/bin/sh

./server_exec &
envoy -c ./sidecar_envoy.yaml
