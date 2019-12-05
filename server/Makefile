PROTOS=contracts/*.pb*
BINDIR=bin
GOOGLEAPIS_DIR=../googleapis

get:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc

all:
	export PATH=$$PATH:~/go/bin; \
	protoc -I ${GOOGLEAPIS_DIR} -I contracts/ --go_out=plugins=grpc:. person.proto; \
	protoc -I ${GOOGLEAPIS_DIR} --include_imports -I contracts/ --include_source_info --descriptor_set_out=contracts/person.pb person.proto

	go build -o $(BINDIR)/personserver server/main.go 

clean:
	rm -rf $(BINDIR)/* $(PROTOS)
