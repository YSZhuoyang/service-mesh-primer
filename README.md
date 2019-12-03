# API-Gateway-Primer

This is a primer project developing a small api gateway using envoy and gRPC.

### Build

1. Pull submodules for the googleapis:

    `git submodule update --recursive`

2. Generate protobuf descriptor file:

    `cd server/Contracts && protoc -I ./googleapis -I. --include_imports --include_source_info --descriptor_set_out=greet.pb *.proto`

3. Build server source project:

    `cd server && dotnet build`

4. Launch services with docker-compose:

    `docker-compose up --build`
