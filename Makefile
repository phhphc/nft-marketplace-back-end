-include .env
export

GO=go
GORUN=$(GO) run
GOBUILD=CGO_ENABLED=0 $(GO) build

http:
	$(GORUN) ./cmd/marketplace

watcher:
	$(GORUN) ./cmd/chain-watcher

puller:
	$(GORUN) ./cmd/data-puller

generate:
	$(GO) generate ./...

build:
	$(GOBUILD) -o ./bin/ ./cmd/*

.PHONY: http generate build 
