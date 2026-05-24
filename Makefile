.PHONY: fmt lint test build all clean

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test -v ./...         


test-race:
	CGO_ENABLED=1 go test -race -v ./...

build:
	mkdir -p bin
	go build -o bin/app.exe ./cmd/app   # .exe на Windows

all: fmt lint test build

clean:
	rm -rf bin