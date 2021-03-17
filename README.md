# Service Mesh Primer

A demo to bootstrap a tiny service mesh with istio which supports:

- Transcoding HTTP+JSON into gRPC+Protobuf
- Server push, streaming data to clients
- Handling Http/1.1, Http/2 and gRPC
- Securing services with Istio sidecar [mTLS](https://istio.io/latest/docs/concepts/security/#mutual-tls-authentication)

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

1. Pull googleapis submodules for gRPC route annotation:

       git submodule update --init --recursive --remote --merge

2. Build source code and docker images (make sure Docker is running):

       make all

## Deploy locally on Kubernetes with [Istio](https://istio.io/)

1. Install Istio:

       curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.9.1 sh -
       cd istio-1.9.1 && export PATH=$PWD/bin:$PATH

2. Generate contract descriptor mounted to istio envoy sidecars (for gRPC transcoding):

       cd deploy/contracts && kubectl create configmap proto-descriptor --from-file=desc.pb

3. Launch istio & services:

       istioctl install --set profile=demo --set values.meshConfig.accessLogFile="/dev/stdout"
       kubectl label namespace default istio-injection=enabled
       kubectl apply -f ./deploy/kube

## Test

- Test Http/1.1 & Http/2 with Curl

      chmod u+x ./test.sh && ./test.sh

- Test Grafana & Jaeger dashboard:

      for i in `seq 1 250`; do ./test.sh; done; # The default Jaeger sampling rate is 1%
      istioctl dashboard jaeger
      istioctl dashboard grafana

- Test server streaming with Web Client

  - Cd to `web-client/public` dir, open `index.html` and click `Get Live Data`.
