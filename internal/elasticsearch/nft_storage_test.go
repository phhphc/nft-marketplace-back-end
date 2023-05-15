package elasticsearch

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"testing"
	"time"
)

var nftStorage *NFTStorage

func loadNFTStorage() {
	storage, err := NewNFTStorage(elasticClient, true)
	if err != nil {
		log.Fatalf("error creating storage: %v", err)
	}
	nftStorage = storage
}

func TestNewNFTStorage(t *testing.T) {
	storage, err := NewNFTStorage(elasticClient, true)
	if err != nil {
		t.Errorf("error creating storage: %v", err)
	}
	if storage == nil {
		t.Errorf("storage is nil")
	}
	nftStorage = storage
}

// func (n *NFTStorage) Insert(ctx context.Context, nft IndexedNFT) error
func TestNftStorage_Insert(t *testing.T) {
	loadNFTStorage()
	nft := generateRandomNFT()

	err := nftStorage.Insert(context.Background(), nft)
	if err != nil {
		t.Errorf("error inserting nft: %v", err)
	}

	insertedNft, err := nftStorage.FindOne(context.Background(), nft.Token, nft.Identifier)

	if err != nil {
		t.Errorf("error finding nft: %v", err)
	}

	require.NotNil(t, insertedNft)
	require.Equal(t, nft.Token, insertedNft.Token)
	require.Equal(t, nft.Identifier, insertedNft.Identifier)
	require.Equal(t, nft.Owner, insertedNft.Owner)
}

func TestNftStorage_Delete(t *testing.T) {
	loadNFTStorage()
	nft := generateRandomNFT()

	err := nftStorage.Insert(context.Background(), nft)
	if err != nil {
		t.Errorf("error inserting nft: %v", err)
	}

	err = nftStorage.Delete(context.Background(), nft.Token, nft.Identifier)
	if err != nil {
		t.Errorf("error deleting nft: %v", err)
	}

	_, err = nftStorage.FindOne(context.Background(), nft.Token, nft.Identifier)
	assert.Nil(t, err)
}

func TestNFTStorage_FullTextSearch(t *testing.T) {
	loadNFTStorage()
	// insert 50 nft
	for i := 0; i < 50; i++ {
		nft := generateRandomNFT()
		err := nftStorage.Insert(context.Background(), nft)
		if err != nil {
			t.Errorf("error inserting nft: %v", err)
		}
	}
}

func generateRandomNFT() IndexedNFT {
	rand.Seed(time.Now().UnixNano())
	tokenBytes := make([]byte, 16)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		panic(err)
	}
	token := hex.EncodeToString(tokenBytes)
	identifier := fmt.Sprintf("%d", rand.Intn(1000000))

	rand.Seed(time.Now().UnixNano())
	ownerBytes := make([]byte, 16)
	if _, err := rand.Read(ownerBytes); err != nil {
		panic(err)
	}
	owner := hex.EncodeToString(ownerBytes)

	metadata := IndexedNFTMetadata{
		Name:        fmt.Sprintf("NFT %s", gofakeit.Name()),
		Description: fmt.Sprintf("Description for NFT %s", gofakeit.Paragraph(1, 3, 30, " ")),
		Image:       fmt.Sprintf("https://example.com/nft/%s/%s/image.jpg", token, identifier),
	}

	return IndexedNFT{
		Token:      token,
		Identifier: identifier,
		Owner:      owner,
		Metadata:   metadata,
	}
}
