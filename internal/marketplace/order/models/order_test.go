package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"testing"
)

func mockOrder() Order {

	startTime, _ := hexutil.DecodeBig("0x0")
	endTime, _ := hexutil.DecodeBig("0xff00000000000000000000000000")
	counter, _ := hexutil.DecodeBig("0x0")

	return Order{
		OrderHash:   "0xc647442747237394433a4ff4bccba6f36d27b1044a32f618ec8fea57ca82cdb1",
		Offerer:     common.HexToAddress("0xA85c072a57bEfE1A907356673137B77ec9b5C985"),
		Signature:   "0xb9edd47357555b24155934bb975365b42b6dad2d6d873c6d0df6e233ff4dbf3f1cb2cb0417ca1e3e6a4b1f3b87fdfba312ba5097f8d372441dc926bfe5fa0e231b",
		OrderType:   big.NewInt(0),
		StartTime:   startTime,
		EndTime:     endTime,
		Counter:     counter,
		Salt:        "0xcf4e492146bbcee31255a79967ae8890fbcd7a7cf8e949e95ae1f84255fd396d",
		IsCancelled: false,
		IsValidated: false,
		Zone:        common.Address{},
		ZoneHash:    "",
	}
}

func TestOrder_GetOrderHash(t *testing.T) {
	order := mockOrder()

	got := order.GetOrderHash()
	want := "0xc647442747237394433a4ff4bccb6f36d27b1044a32f618ec8fea57ca82cdb1"

	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}
