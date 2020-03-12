
get:
	wget https://github.com/grpc/grpc-web/releases/download/1.0.7/protoc-gen-grpc-web-1.0.7-darwin-x86_64 --no-check-certificate
	mv protoc-gen-grpc-web-1.0.7-darwin-x86_64 /usr/local/bin/protoc-gen-grpc-web
	chmod +x /usr/local/bin/protoc-gen-grpc-web

gen:
	export PATH=$$PATH:/usr/local/bin/protoc-gen-grpc-web; \
	cd contracts; \
	protoc -I . -I ./googleapis ./greeter-go/*.proto \
	--js_out=import_style=commonjs:../src \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:../src
	# Please remove imports for annotation which is unused from generated files (waiting for official fix)

clean:
	rm -f ./public/*_pb.js
