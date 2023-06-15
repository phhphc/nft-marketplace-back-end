package entities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Profile struct {
	Username  string         `json:"username"`
	Address   common.Address `json:"address"`
	Metadata  map[string]any `json:"metadata"`
	Signature []byte         `json:"signature"`
}

type EncodeProfile struct {
	Username string         `json:"username"`
	Address  string         `json:"address"`
	Metadata map[string]any `json:"metadata"`
}

func (p *Profile) Verify() bool {
	// Get the public key from the signature
	sigPublicKeyECDSA, _ := crypto.SigToPub(p.Hash(), p.Signature)

	// Convert the public key to an Ethereum address
	publicAddress := crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex()

	return p.Address.Hex() == publicAddress
}

func (p *Profile) Hash() []byte {
	enc := p.Encode()
	hash := crypto.Keccak256Hash(enc)
	return hash.Bytes()
}

func (p *Profile) Encode() []byte {
	var _ bytes.Buffer

	encodeProfile := EncodeProfile{
		Username: p.Username,
		Address:  p.Address.Hex(),
		Metadata: p.Metadata,
	}

	encoded, err := json.Marshal(encodeProfile)
	fmt.Println(encoded)
	if err != nil {
		return []byte{}
	}

	return encoded
}
