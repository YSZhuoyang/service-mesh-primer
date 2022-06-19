
PLATFORM=darwin-x86_64
GRPC_WEB_VERSION=1.3.1

ifeq ($(OS), DEBIAN)
	PLATFORM=linux-x86_64
endif

get:
	wget https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-$(PLATFORM) --no-check-certificate
	mv protoc-gen-grpc-web-${GRPC_WEB_VERSION}-$(PLATFORM) ./protoc-gen-grpc-web
	chmod +x ./protoc-gen-grpc-web

gen:
	export PATH=$$PATH:$$PWD; \
	protoc -I ../contracts -I ../contracts/googleapis ../contracts/greeter-go/*.proto \
		--js_out=import_style=commonjs:./src \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src

	# Remove imports for annotation which is unused from generated files (waiting for official fix)
ifeq ($(OS), DEBIAN)
	ls -1 ./src/greeter-go/*.js | xargs sed -i '/annotations_pb/d'
else
	ls -1 ./src/greeter-go/*.js | xargs sed -i "" '/annotations_pb/d'
endif

clean:
	rm -f ./src/**/*_pb.js
