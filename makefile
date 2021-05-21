
services:
	cd contracts && make build
	cd go-service && make all && docker build -t go-service .
	cp -r contracts ./dotnet-service && cd dotnet-service && docker build -t dotnet-service . && rm -r contracts

client:
	cd web-client && \
	make get && \
	make gen && \
	yarn build

all:
	make services
	make client
