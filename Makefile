-include .env
export

GO=go
GORUN=$(GO) run
GOBUILD=$(GO) build

http:
	$(GORUN) ./cmd/marketplace

generate:
	$(GO) generate ./...

build:
	$(GOBUILD) -o ./bin/ ./cmd/marketplace

.PHONY: http generate build 