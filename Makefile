BINARY_NAME=main
 
build:
	go mod tidy
	go build -o build/${BINARY_NAME} cmd/main.go
 
clean:
	go clean
	rm -rf build

codegen:
	protoc --go-grpc_out=. --go_out=. --proto_path=./proto ./proto/*.proto google/protobuf/timestamp.proto google/protobuf/empty.proto 
