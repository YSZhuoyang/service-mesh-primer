# API-Gateway-Primer

Bootstrapping a tiny api gateway using envoy transcoding HTTP+JSON into gRPC+Protobuf.

                   gateway           sidecar proxy  service
        http json   |--|  grpc protobuf  |--|        |--|
       ------------>|  |---------------->|  |------->|  |
                    |--|                 |--|        |--|

## Build

1. Pull submodules for the googleapis:

    `git submodule update --recursive`

2. Generate api source code and build:

    `make all`

3. Launch services locally (envoy must be installed):

    `chmod u+x ./launch_local.sh`
    `./launch_local.sh`

## Test

    `cd test && chmod u+x ./test.sh && ./test.sh`
