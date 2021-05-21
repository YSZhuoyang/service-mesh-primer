# Service Mesh Primer

A demo to bootstrap a tiny service mesh with istio which supports:

- Transcoding HTTP+JSON into gRPC+Protobuf
- Server push, streaming data to clients
- Handling Http/1.1, Http/2 and gRPC
- Securing services with Istio sidecar [mTLS](https://istio.io/latest/docs/concepts/security/#mutual-tls-authentication)

                     gateway             sidecar     service1
      http(JSON)/grpc |--|                 |--|  grpc  |--|
      --------------->|  |---------------->|  |------->|  |
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

       curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.10.0 sh -
       cd istio-1.10.0 && export PATH=$PWD/bin:$PATH

2. Generate contract descriptor mounted to istio envoy sidecars (for gRPC transcoding):

       kubectl create configmap proto-descriptor --from-file=contracts/desc.pb

3. Launch istio & services:

       istioctl install -f kube/istio-install/config.yaml -y
       kubectl label namespace default istio-injection=enabled
       kubectl apply -f ./kube/deploy

## Test

- Test Http/1.1 & Http/2 with Curl

      ./test.sh

- Test Grafana & Jaeger dashboard:

      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/addons/jaeger.yaml
      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/addons/prometheus.yaml
      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/addons/grafana.yaml

      for i in `seq 1 250`; do ./test.sh; done; # The default Jaeger sampling rate is 1%

      istioctl dashboard jaeger
      istioctl dashboard grafana

- Test server streaming with Web Client

  - Cd to `web-client/public` dir, open `index.html` and click `Get Live Data`.
