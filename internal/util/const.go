package util

import "github.com/ethereum/go-ethereum/common"

var (
	ZeroAddress common.Address
)

func init() {
	ZeroAddress = common.Address{}
}
