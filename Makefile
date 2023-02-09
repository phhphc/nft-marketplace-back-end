-include .env
export

GO=go
GORUN=$(GO) run
GOBUILD=CGO_ENABLED=0 $(GO) build

http:
	$(GORUN) ./cmd/marketplace

watcher:
	$(GORUN) ./cmd/chain-watcher

generate:
	$(GO) generate ./...

build:
	$(GOBUILD) -o ./bin/ ./cmd/marketplace
	$(GOBUILD) -o ./bin/ ./cmd/chain-watcher

.PHONY: http generate build 
