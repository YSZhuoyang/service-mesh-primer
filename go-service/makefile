install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get:
	go get -u

build:
	mkdir -p rpc/dotnet-service rpc/go-service
	export PATH=$$PATH:$$(go env GOPATH)/bin; \
	protoc -I ../contracts/greeter-go -I ../contracts/googleapis \
		--go_out=rpc/go-service --go_opt=paths=source_relative --go-grpc_out=rpc/go-service --go-grpc_opt=paths=source_relative ../contracts/greeter-go/*.proto
	protoc -I ../contracts/greeter-dotnet -I ../contracts/googleapis \
		--go_out=rpc/dotnet-service --go_opt=paths=source_relative --go-grpc_out=rpc/dotnet-service --go-grpc_opt=paths=source_relative ../contracts/greeter-dotnet/*.proto
	go build -o server_exec
	go mod tidy

build-frozen:
	go build -mod=readonly -o server_exec

all:
	make install
	make build

clean:
	rm -f server_exec
	rm -f -r rpc
