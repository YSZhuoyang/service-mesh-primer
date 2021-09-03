FROM golang:1.17
WORKDIR /go/server

# Build src
COPY ./ ./
RUN make build-frozen

ENTRYPOINT ./server_exec
