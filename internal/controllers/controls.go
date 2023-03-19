package controllers

import (
	"github.com/phhphc/nft-marketplace-back-end/internal/services"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type Controls struct {
	lg      *log.Logger
	service services.Servicer
}

func New(service services.Servicer) *Controls {
	return &Controls{
		lg:      log.GetLogger(),
		service: service,
	}
}
