BINARY_NAME=main
 
build:
	make clean
	go mod tidy
	make grpc
	go build -o build/${BINARY_NAME} cmd/main.go
 
clean:
	rm -rf build
	go clean

grpc:
	protoc --go-grpc_out=. --go_out=. --proto_path=./proto ./proto/*.proto google/protobuf/timestamp.proto google/protobuf/empty.proto 
