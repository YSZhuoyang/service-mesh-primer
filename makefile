
services:
	cd contracts && make build
	cd go-service && make all && docker build -t go-service .
	cd dotnet-service && dotnet build && docker build -t dotnet-service .

client:
	cd web-client && \
	make get && \
	make gen && \
	yarn build

all:
	make services
	make client
