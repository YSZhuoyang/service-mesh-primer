# API-Gateway-Primer

Bootstraping a tiny api gateway using envoy transcoding HTTP+JSON into gRPC+Protobuf.

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

    `chmod u+x ./go-service/start_service.sh`
    `cd ./go-service && ./start_service.sh`

    `chmod u+x ./dotnet-service/start_service.sh`
    `cd ./dotnet-service && ./start_service.sh`

    `chmod u+x ./gateway/start_service.sh`
    `cd ./gateway && ./start_service.sh`

## Test

    `cd test && chmod u+x ./test.sh && ./test.sh`
