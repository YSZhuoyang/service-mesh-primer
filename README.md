# API-Gateway-Primer

Bootstrapping a tiny API gateway using envoy transcoding HTTP+JSON into gRPC+Protobuf.

                   gateway           sidecar proxy  service1
        http json   |--|       grpc      |--|        |--|
       ------------>|  |---------------->|  |------->|  |
                    |--|        |        |--|        |--|
                                |
                                |    sidecar proxy  service2
                                |        |--|        |--|
                                -------->|  |------->|  |
                                         |--|        |--|

## Build & Run

1. Pull googleapis submodules for gRPC route annotation:

    `git submodule update --recursive --remote --merge`

2. Generate contract source code and build:

    `make all`

3. Launch services:

    `docker-compose up --build`

## Test

    `chmod u+x ./test.sh && ./test.sh`
