all: build test

test: build
	@go test ./...

build:
	@go fmt ./... && go vet ./... && go build -buildmode=plugin

clean:
	@go clean
