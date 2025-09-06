
GRPC_WEB_PLATFORM=darwin-aarch64
GRPC_WEB_VERSION=2.0.1
JS_PROTOC_PLATFORM=osx-aarch_64
JS_PROTOC_VERSION=4.0.0


ifeq (${OS}, DEBIAN)
	GRPC_WEB_PLATFORM=linux-aarch64
	JS_PROTOC_PLATFORM=linux-aarch_64
endif

get-protoc-gen-js:
	wget https://github.com/protocolbuffers/protobuf-javascript/releases/download/v${JS_PROTOC_VERSION}/protobuf-javascript-${JS_PROTOC_VERSION}-${JS_PROTOC_PLATFORM}.zip --no-check-certificate
	unzip protobuf-javascript-${JS_PROTOC_VERSION}-${JS_PROTOC_PLATFORM}.zip -d ./protoc-js
	mv -f ./protoc-js/protobuf-javascript-${JS_PROTOC_VERSION}-${JS_PROTOC_PLATFORM}/bin/protoc-gen-js .
	rm -r -f ./protoc-js
	rm -f protobuf-javascript-${JS_PROTOC_VERSION}-${JS_PROTOC_PLATFORM}.zip

get-protoc-gen-web:
	wget https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-${GRPC_WEB_PLATFORM} --no-check-certificate
	mv protoc-gen-grpc-web-${GRPC_WEB_VERSION}-${GRPC_WEB_PLATFORM} ./protoc-gen-grpc-web
	chmod +x ./protoc-gen-grpc-web

get:
	make get-protoc-gen-js
	make get-protoc-gen-web

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
