# Service-Mesh-Primer

Bootstrapping a tiny service mesh with istio which supports:

- Transcoding HTTP+JSON into gRPC+Protobuf
- Server push, streaming data to clients
- Handling Http/1.1, Http/2 and gRPC

                   gateway             sidecar     service1
        http json   |--|                 |--|  grpc  |--|
       ------------>|  |---------------->|  |------->|  |
                    |--|        |        |--|        |--|
                                |
                                |      sidecar     service2
                                |        |--|  grpc  |--|
                                -------->|  |------->|  |
                                         |--|        |--|

## Build

   `make all`

## Deploy locally on Kubernetes

1. Pull googleapis submodules for gRPC route annotation:

    `git submodule update --init --recursive --remote --merge`

2. Generate contract descriptor which can be mounted to istio envoy sidecars:

    `cd contracts && make build`
    `kubectl create configmap proto-descriptor --from-file=desc.pb`

3. Launch istio & services:

    `istioctl manifest apply --set values.global.proxy.accessLogFile="/dev/stdout"`
    `kubectl label namespace default istio-injection=enabled`
    `kubectl apply -f ./kube`

## Test

- Test Http/1.1 & Http/2 with Curl

      chmod u+x ./test.sh && ./test.sh

- Test server streaming with Web Client

  - Cd to `web-client/public` dir, open `index.html` and click `Get Live Data`.
