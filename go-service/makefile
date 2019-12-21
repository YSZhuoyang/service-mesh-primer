get:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go

build:
	export PATH=$$PATH:~/go/bin; \
	cd contracts; \
	protoc -I . -I ./googleapis --go_out=plugins=grpc:. ./greeter-go/*.proto
	go build -o server_exec

all:
	make get
	make build

clean:
	rm -f server_exec; \
	rm -f ./contracts/*/*.pb*
