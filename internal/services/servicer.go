package services

type Servicer interface {
	OrderService
	NftNewService
}

var _ Servicer = (*Services)(nil)
