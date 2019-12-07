# API-Gateway-Primer

Bootstrapping a tiny api gateway using envoy transcoding HTTP+JSON into gRPC+Protobuf.

                   gateway           sidecar proxy  service
        http json   |--|       grpc      |--|        |--|
       ------------>|  |---------------->|  |------->|  |
                    |--|        |        |--|        |--|
                                |
                                |
                                |        |--|        |--|
                                -------->|  |------->|  |
                                         |--|        |--|

## Build

1. Pull submodules for the googleapis:

    `git submodule update --recursive`

2. Generate api source code and build:

    `make all`

3. Launch services locally (envoy must be installed):

    `docker-compose up --build`

## Test

    `cd test && chmod u+x ./test.sh && ./test.sh`
