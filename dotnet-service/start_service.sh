#!/bin/sh

dotnet run &
envoy -c ./sidecar_envoy.yaml
