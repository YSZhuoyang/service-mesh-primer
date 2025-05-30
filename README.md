# Service Mesh Primer

[![Open in Visual Studio Code](https://open.vscode.dev/badges/open-in-vscode.svg)](https://open.vscode.dev/YSZhuoyang/service-mesh-primer)

A demo to bootstrap a tiny service mesh with istio which supports:

- Transcoding HTTP+JSON into gRPC+Protobuf
- Server push, streaming data to clients
- Handling Http/1.1, Http/2 and gRPC
- Securing services with Istio sidecar [mTLS](https://istio.io/latest/docs/concepts/security/#mutual-tls-authentication)

                     gateway     waypoint         service1
      http(JSON)/grpc +--+         +--+       grpc  +--+
      --------------->|  |-------->|  |------------>|  |
                      +--+         +--+    |        +--+
                                           |
                                           |      service2
                                           |  grpc  +--+
                                           |------->|  |
                                                    +--+

## Build

1. Pull googleapis submodules for gRPC route annotation:

       git submodule update --init --recursive --remote --merge

2. Build source code and docker images (make sure Docker is running):

       make all

## Deploy locally on Kubernetes with [Istio](https://istio.io/)

1. Install Istio:

       curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.26.1 sh -
       cd istio-1.26.1 && export PATH=$PWD/bin:$PATH

2. Generate contract descriptor mounted to istio envoy gateway (for gRPC transcoding):

       kubectl create namespace istio-system
       kubectl create configmap proto-descriptor --from-file=contracts/desc.pb -n istio-system

3. Install kube gateway api:

       kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.3.0/standard-install.yaml
       <!-- kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.3.0/experimental-install.yaml -->

4. Launch istio & services:

       istioctl install -f kube/istio/istio-operator.yaml -y
       <!-- Add all services in default namespace into ambient mesh -->
       kubectl label namespace default istio.io/dataplane-mode=ambient
       <!-- Enroll all services in default namespace to use a waypoint, any requests using the ambient data plane mode, to any service running in this namespace, will be routed through the waypoint for L7 processing and policy enforcement -->
       <!-- https://istio.io/latest/docs/ambient/usage/waypoint/#useawaypoint -->
       istioctl waypoint apply -n default --enroll-namespace
       kubectl apply -f ./kube/services

## Test

- Test Http/1.1 & Http/2 requests:

      ./test.sh

- Test Grafana & Jaeger dashboard:

      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.26/samples/addons/jaeger.yaml
      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.26/samples/addons/prometheus.yaml
      kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.26/samples/addons/grafana.yaml

      for i in `seq 1 250`; do ./test.sh; done; # The default Jaeger sampling rate is 1%

      istioctl dashboard jaeger
      istioctl dashboard grafana

- Test server streaming with a web client:

      cd web-client/public
      # open `index.html` in a browser and click `Get Live Data` button.

- Cleanup

      kubectl label namespace default istio.io/use-waypoint-
      istioctl waypoint delete --all
      kubectl label namespace default istio.io/dataplane-mode-
      istioctl uninstall -y --purge
      kubectl delete namespace istio-system
      kubectl delete -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.3.0/standard-install.yaml
      <!-- kubectl delete -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.3.0/experimental-install.yaml -->
      kubectl delete deployment --all
      kubectl delete svc dotnet-service go-service
      kubectl delete configmap proto-descriptor
