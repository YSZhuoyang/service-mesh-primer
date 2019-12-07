#!/bin/bash

curl POST -d '{"msg": "Hello service"}' -H "Content-Type: application/json" http://0.0.0.0:80/greet/hello
