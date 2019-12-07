#!/bin/bash

curl -X POST -d '{"Msg": "Hello service"}' -H "Content-Type:application/json" http://0.0.0.0:80/greet/hello
