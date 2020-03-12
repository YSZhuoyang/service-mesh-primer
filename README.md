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

    `git submodule update --init --recursive --remote --merge`

2. Generate contract source code and build:

    `make all`

3. Launch services:

    `docker-compose up --build`

## Test

### Test Http/1.1 & Http/2 with Curl

    `chmod u+x ./test.sh && ./test.sh`

### Test server streaming with Web Client

    `cd web-client`
    `make get`
    `make gen` # And follow instructions to remove the annotation imports from generated files
    `yarn build`

    Cd to `/public` folder, open `index.html` and click `Get Live Data`.
