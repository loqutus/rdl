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

test:
	./bin/rdl-linux torrent test/cosmos-laundromat.torrent gdrive6:/

clean:
	go clean
	rm bin/${BINARY_NAME}-linux
	rclone purge gdrive6:/Cosmos\ Laundromat || true

default: get build test clean