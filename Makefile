include .env
export

GO=go
GORUN=$(GO) run

http:
	$(GORUN) ./cmd/marketplace

generate:
	$(GO) generate ./...

.PHONY: http generate