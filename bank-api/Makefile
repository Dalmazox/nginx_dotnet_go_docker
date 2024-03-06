export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GOPROXY=https://proxy.golang.org,direct

all: clean build
.PHONY: all

build:
	go build -o ./dist/bank-api ./cmd/server.go
.PHONY: build

run:
	go run ./cmd/server.go
.PHONY: run

clean: tidy
	rm -rf ./dist
.PHONY: clean

tidy:
	go mod tidy
.PHONY: tidy