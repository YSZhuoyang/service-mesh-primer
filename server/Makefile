PROTOS=sample/msglib/*.pb*
BINDIR=bin
GOOGLEAPIS_DIR=../googleapis

get:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc

all:
	export PATH=$$PATH:~/go/bin; \
	protoc -I${GOOGLEAPIS_DIR} -Isample/msglib --go_out=plugins=grpc:. person.proto; \
	protoc -I${GOOGLEAPIS_DIR} --include_imports -Isample/msglib --include_source_info --descriptor_set_out=sample/msglib/person.pb person.proto

	go build -o $(BINDIR)/personserver server/main.go 
	go build -o $(BINDIR)/personclient client/main.go 

clean:
	rm -rf $(BINDIR)/* $(PROTOS)
