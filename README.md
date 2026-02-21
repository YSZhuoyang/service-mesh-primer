# Service Mesh Primer

[![Open in Visual Studio Code](https://img.shields.io/badge/Open%20in-VS%20Code-007ACC?logo=visual-studio-code)](https://vscode.dev/github/YSZhuoyang/service-mesh-primer)

A demo to bootstrap a tiny service mesh with istio which supports:

- Transcoding HTTP+JSON into gRPC+Protobuf
- Server push, streaming data to clients
- Handling Http/1.1, Http/2 and gRPC
- Securing services with [mTLS](https://istio.io/latest/docs/concepts/security/#mutual-tls-authentication)

                      gateway     waypoint       service1
      http(JSON)/grpc  +--+   grpc   +--+    grpc  +--+
      ---------------->|  |--------->|  |--------->|  |
                       +--+          +--+    |     +--+
                                             |
                                             |   service2
                                             |     +--+
                                             ----->|  |
                                             grpc  +--+

## Build

1. Pull googleapis submodules for gRPC route annotation:

       git submodule update --init --recursive --remote --merge

2. Build source code and docker images (make sure Docker is running):

       make all

## Deploy locally on Kubernetes with [Istio](https://istio.io/)

1. Optional: Download Istio (skip if using dev container):

       curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.29.0 sh -
       cd istio-1.29.0 && export PATH=$PWD/bin:$PATH

2. Generate contract descriptor mounted to istio envoy gateway (for gRPC transcoding):

       kubectl create configmap proto-descriptor --from-file=contracts/desc.pb

3. Install kube gateway api:

       <!-- Run below with flag if size exceeds limit: --server-side -->
       <!-- kubectl apply --server-side -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.4.1/standard-install.yaml -->
       kubectl apply --server-side -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.4.1/experimental-install.yaml

4. Launch istio & services:

       kubectl create namespace istio-system
       istioctl install -f kube/istio/istio-operator.yaml -y
       <!-- Add all services in default namespace into ambient mesh -->
       kubectl label namespace default istio.io/dataplane-mode=ambient
       <!-- Enroll all services in default namespace to use a waypoint, any requests using the ambient data plane mode, to any service running in this namespace, will be routed through the waypoint for L7 processing and policy enforcement -->
       <!-- https://istio.io/latest/docs/ambient/usage/waypoint/#useawaypoint -->
       istioctl waypoint apply -n default --enroll-namespace
       kubectl apply --server-side -f ./kube/services-stage-1
       kubectl apply --server-side -f ./kube/services-stage-2
       <!-- Use tunnel to access API gateway e.g.: -->
       kubectl port-forward svc/istio-ingressgateway-istio 8080:80

## Test

- Test Http/1.1 & Http/2 requests:

      ./test.sh

- Test Grafana & Jaeger dashboard:

      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.29/samples/addons/jaeger.yaml
      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.29/samples/addons/prometheus.yaml
      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.29/samples/addons/grafana.yaml

      for i in `seq 1 250`; do ./test.sh; done; # The default Jaeger sampling rate is 1%

      istioctl dashboard jaeger
      istioctl dashboard grafana

- Test server streaming with a web client:

      # In Dev Container, run:
      cd web-client && python3 -m http.server 3000 --directory public
      # Otherwise run:
      cd web-client/public
      # open `index.html` in a browser and click `Get Live Data` button.

- Cleanup

      kubectl label namespace default istio.io/use-waypoint-
      istioctl waypoint delete --all
      kubectl label namespace default istio.io/dataplane-mode-
      istioctl uninstall -y --purge
      kubectl delete namespace istio-system
      <!-- kubectl delete -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.4.1/standard-install.yaml -->
      kubectl delete -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.4.1/experimental-install.yaml
      kubectl delete deployment --all
      kubectl delete svc dotnet-service go-service
      kubectl delete configmap proto-descriptor
