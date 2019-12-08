#!/bin/sh

dotnet run &
envoy -c ./sidecar-envoy.yaml
