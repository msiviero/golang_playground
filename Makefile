BINARY_NAME=main

.PHONY: build

build:
	make clean
	make code_gen
	make deps
	go build -o build/${BINARY_NAME} cmd/main.go cmd/wire_gen.go
 
deps:
	go mod tidy

clean:
	rm -rf build
	go clean

code_gen:
	rm -rf ./internal/grpc
	wire cmd/wire.go
	protoc --go-grpc_out=. --go_out=. --proto_path=./proto ./proto/*.proto
