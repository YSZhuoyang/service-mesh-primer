FROM golang:1.22
WORKDIR /go/server

# Build src
COPY ./ ./
RUN make build-frozen

ENTRYPOINT ./server_exec
