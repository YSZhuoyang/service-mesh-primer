# API-Gateway-Primer

Bootstrapping a tiny API gateway using envoy which supports:

- Transcoding HTTP+JSON into gRPC+Protobuf
- Streaming data to clients
- Handling Http/1.1, Http/2 and gRPC

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

* Test Http/1.1 & Http/2 with Curl

      chmod u+x ./test.sh && ./test.sh

* Test server streaming with Web Client

  - Cd to `web-client/public` dir, open `index.html` and click `Get Live Data`.
