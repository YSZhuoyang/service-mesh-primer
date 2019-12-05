
all:
	cd gateway && make build
	cd go-service && make all
