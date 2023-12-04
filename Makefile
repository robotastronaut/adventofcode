test:
	go test ./...

lint:
	golangci-lint run ./...

check: lint test

.PHONY: test lint check
