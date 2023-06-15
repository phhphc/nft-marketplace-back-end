package util

import "github.com/ethereum/go-ethereum/common"

var (
	ZeroAddress common.Address
	ZeroHash    common.Hash

	TruePointer  *bool
	FalsePointer *bool
)

func init() {
	ZeroAddress = common.Address{}
	ZeroHash = common.Hash{}

	TruePointer = &[]bool{true}[0]
	FalsePointer = &[]bool{false}[0]
}
