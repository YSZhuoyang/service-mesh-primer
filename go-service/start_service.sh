#!/bin/sh

./server_exec &
envoy -c ./sidecar-envoy.yaml
