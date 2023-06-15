package entities

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func CreateWallet() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return privateKey, publicKeyECDSA
}

func AddressFromPublicKey(publicKey *ecdsa.PublicKey) common.Address {
	return crypto.PubkeyToAddress(*publicKey)
}

func TestHash(t *testing.T) {
	metadata := map[string]any{
		"bio":        "This is my template bio. Beginning of new era nft.",
		"image_url":  "http://example.com",
		"banner_url": "http://i.sean/example.png",
		"email":      "truongphuhung@example.com",
	}

	profile := Profile{
		Username: "tph",
		Metadata: metadata,
		Address:  common.HexToAddress("0x96216849c49358B10257cb55b28eA603c874b05E"), // go-ethereum address
	}
	hash := profile.Hash()

	assert.NotEmpty(t, hash, "should not be empty")
	assert.IsType(t, []byte{}, hash, "should be a array of bytes")
}

//func TestEncode(t *testing.T) {
//	metadata := map[string]any{
//		"bio":        "This is my template bio. Beginning of new era nft.",
//		"image_url":  "http://example.com",
//		"banner_url": "http://i.sean/example.png",
//		"email":      "truongphuhung@example.com",
//	}
//
//	profile := Profile{
//		Username: "tph",
//		Metadata: metadata,
//		Address:  common.HexToAddress("0x96216849c49358B10257cb55b28eA603c874b05E"), // go-ethereum address
//	}
//
//	profileJSON := `{"username":"tph","address":"0x96216849c49358B10257cb55b28eA603c874b05E","metadata":{"bio":"This is my template bio. Beginning of new era nft.","image_url":"http://example.com","banner_url":"http://i.sean/example.png","email":"truongphuhung@example.com"}}`
//
//	encodeProfile := profile.Encode()
//	expected := []byte(profileJSON)
//	assert.NotEmpty(t, encodeProfile, "should not be empty")
//	assert.Equal(t, expected, encodeProfile, "should be equal")
//}

func TestVerify(t *testing.T) {
	privateKey, publicKey := CreateWallet()
	address := AddressFromPublicKey(publicKey)

	metadata := map[string]any{
		"bio":        "This is my template bio. Beginning of new era nft.",
		"image_url":  "http://example.com",
		"banner_url": "http://i.sean/example.png",
		"email":      "truongphuhung@example.com",
	}

	profile := Profile{
		Username: "tph",
		Metadata: metadata,
		Address:  address, // go-ethereum address
	}

	// A test case verify the signature of the profile after signed
	t.Run("Verify the signature", func(t *testing.T) {
		signature, err := crypto.Sign(profile.Hash(), privateKey)
		if err != nil {
			t.Errorf("should create signature, got: %v", err)
		}
		profile.Signature = signature
		isValid := profile.Verify()
		assert.True(t, isValid, "should be valid, got: %v", isValid)
	})
}
