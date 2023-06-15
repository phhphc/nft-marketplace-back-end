package util

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"strings"
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

func MustMapJsonToBytes(
	m map[string]any,
) (bs []byte) {
	if m == nil {
		return
	}

	bs, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return
}

func MustHexToBytes(
	s string,
) (bs []byte) {
	bs, err := hex.DecodeString(strings.TrimLeft(s, "0x"))
	if err != nil {
		panic(err)
	}
	return
}

func BytesToHex(
	bs []byte,
) (s string) {
	s = "0x" + hex.EncodeToString(bs)
	return
}
