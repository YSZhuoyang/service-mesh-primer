#!/bin/sh
./bin/personserver &
envoy -c ./conf/envoy.yaml
