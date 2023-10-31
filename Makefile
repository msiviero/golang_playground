.PHONY: build

build:
	make clean
	make code_gen
	make deps
	go build -o build/server ./cmd
 
deps:
	go mod tidy

clean:
	rm -rf build
	go clean

code_gen:
	rm -rf ./internal/grpc
	wire cmd/wire.go
	protoc --go-grpc_out=. --go_out=. --proto_path=./proto ./proto/*.proto

grpc_client:
	grpcui -plaintext -import-path=/Users/marco.siviero/Downloads/go-playground/proto -proto=user_route.proto 127.0.0.1:50051
