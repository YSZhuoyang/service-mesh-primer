
get:
	wget https://github.com/grpc/grpc-web/releases/download/1.2.1/protoc-gen-grpc-web-1.2.1-darwin-x86_64 --no-check-certificate
	mv protoc-gen-grpc-web-1.2.1-darwin-x86_64 ./protoc-gen-grpc-web
	chmod +x ./protoc-gen-grpc-web

gen:
	export PATH=$$PATH:$$PWD; \
	protoc -I ../contracts -I ../contracts/googleapis ../contracts/greeter-go/*.proto \
	--js_out=import_style=commonjs:./src \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src
	ls -1 ./src/greeter-go/*.js | xargs sed -i "" '/annotations_pb/d'
	# Remove imports for annotation which is unused from generated files (waiting for official fix)

clean:
	rm -f ./src/**/*_pb.js
