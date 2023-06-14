package util

import (
	"encoding/json"
	"errors"
	"math/big"
)

var ErrCastError = errors.New("error cast")

func MustStringToBigInt(
	s string,
) *big.Int {
	b, ok := new(big.Int).SetString(s, 0)
	if !ok {
		panic(ErrCastError)
	}
	return b
}

func MustBytesToMapJson(
	bs []byte,
) (m map[string]any) {
	if len(bs) > 0 {
		m = make(map[string]any)
		err := json.Unmarshal(bs, &m)
		if err != nil {
			panic(err)
		}
	}
	return
}
