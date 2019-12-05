FROM envoyproxy/envoy:latest

ADD . .

RUN apt update && apt -q install -y curl

RUN make build

RUN chmod u+x ./start_service.sh
ENTRYPOINT ./start_service.sh
