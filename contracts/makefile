
build:
	protoc \
	-I ./greeter-go \
	-I ./greeter-dotnet \
	-I ./googleapis \
	--include_imports --include_source_info \
	--descriptor_set_out=./desc.pb ./**/*.proto
