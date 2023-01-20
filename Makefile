include .env
export

GORUN=go run

http:
	$(GORUN) ./cmd/marketplace

.PHONY: http