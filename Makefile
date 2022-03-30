.DEFAULT_GOAL := default
.PHONY: all helm
BINARY_NAME=rdl

get:
	go mod tidy
	go get ./...

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/rdl/main.go

build_darwin:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin cmd/rdl/main.go

run:
	./bin/${BINARY_NAME}-linux

run_darwin:
	./bin/${BINARY_NAME}-darwin

test:
	go test -v -count=1 ./... 

clean:
	go clean
	rm bin/${BINARY_NAME}-linux

default: get build