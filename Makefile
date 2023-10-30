BINARY_NAME=main
 
build:
	go mod tidy
	go build -o build/${BINARY_NAME} cmd/main.go
 
clean:
	go clean
	rm -rf build