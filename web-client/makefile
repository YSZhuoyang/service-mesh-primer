
build:
	cd contracts; \
	protoc -I . -I ./googleapis ./greeter-go/*.proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:../public

clean:
	rm -f ./contracts/*/*.pb*
