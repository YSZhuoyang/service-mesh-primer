
build:
	cd contracts; \
	protoc \
	-I ./greeter \
	-I ./googleapis \
	--include_imports --include_source_info \
	--descriptor_set_out=./desc.pb ./**/*.proto
