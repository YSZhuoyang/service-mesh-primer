# Service Mesh Primer

[![Open in Visual Studio Code](https://open.vscode.dev/badges/open-in-vscode.svg)](https://open.vscode.dev/YSZhuoyang/service-mesh-primer)

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

       curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.12.1 sh -
       cd istio-1.12.1 && export PATH=$PWD/bin:$PATH

2. Generate contract descriptor mounted to istio envoy sidecars (for gRPC transcoding):

       kubectl create configmap proto-descriptor --from-file=contracts/desc.pb

3. Launch istio & services:

       istioctl install -f kube/istio/istio-operator.yaml -y
       kubectl label namespace default istio-injection=enabled
       kubectl apply -f ./kube/services

## Test

- Test Http/1.1 & Http/2 requests:

      ./test.sh

- Test Grafana & Jaeger dashboard:

      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.12/samples/addons/jaeger.yaml
      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.12/samples/addons/prometheus.yaml
      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.12/samples/addons/grafana.yaml

      for i in `seq 1 250`; do ./test.sh; done; # The default Jaeger sampling rate is 1%

      istioctl dashboard jaeger
      istioctl dashboard grafana

- Test server streaming with a web client:

      cd web-client/public
      # open `index.html` in a browser and click `Get Live Data` button.

- Cleanup

      istioctl manifest generate | kubectl delete --ignore-not-found=true -f -
      kubectl delete namespace istio-system
      kubectl label namespace default istio-injection-
      kubectl delete deployment --all
      kubectl delete svc dotnet-service go-service
      kubectl delete configmap proto-descriptor
