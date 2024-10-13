.PHONY: default
default: build lint test

.PHONY: build
build:
	go build ./...

.PHONY: lint
lint:
	go vet ./...
	go fmt ./...

.PHONY: test
test:
	go test ./...
