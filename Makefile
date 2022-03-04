# APP_NAME=

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

.PHONY: build
build:
	go env -w GOPRIVATE=github.com/on0z/*
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(GOBIN)/app ./cmd/${APP_NAME}/main.go

.PHONY: build_linux_amd64
build_linux_amd64:
	go env -w GOPRIVATE=github.com/on0z/*
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GOBIN)/app ./cmd/${APP_NAME}/main.go

.PHONY: build_linux_arm64
build_linux_arm64:
	go env -w GOPRIVATE=github.com/on0z/*
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o $(GOBIN)/app ./cmd/${APP_NAME}/main.go

.PHONY: start
start:
	go build -o $(GOBIN)/app ./cmd/${APP_NAME}/main.go
	./bin/app

.PHONY: test-v
test-v:
	go test -v ./pkg/...

.PHONY: test
test:
	LogDiscard=True go test ./pkg/...

.PHONY: clean
clean:
	go mod tidy
	go clean --modcache
	rm -rf $(GOBIN)/app
