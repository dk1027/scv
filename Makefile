TARGET=scv

all: build test

test: build
	@go test ./...

build:
	@go fmt ./... && go vet ./... && go build -o ${TARGET}
	@pushd watchdir; make

clean:
	@go clean
