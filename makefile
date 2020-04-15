
services:
	cd deploy && make build
	cd go-service && make all
	cd dotnet-service && dotnet build

client:
	cd web-client && \
	make get && \
	make gen && \
	yarn build

all:
	make services
	make client
