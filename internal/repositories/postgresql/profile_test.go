package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/stretchr/testify/require"
	"github.com/tabbed/pqtype"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func createRandomProfile(t *testing.T) Profile {
	metadata := entities.ProfileMetadata{
		Bio:       RandomString(100),
		ImageUrl:  RandomString(50),
		BannerUrl: RandomString(50),
		Email:     RandomString(20),
	}
	profile := entities.Profile{
		Address:   common.HexToAddress(RandomAddress()),
		Username:  RandomString(10),
		Metadata:  metadata,
		Signature: []byte(RandomHex(82))[2:],
	}

	json, _ := json.Marshal(profile.Metadata)
	metadataRaw := pqtype.NullRawMessage{
		RawMessage: json,
		Valid:      true,
	}
	expected := UpsertProfileParams{
		Address:   profile.Address.Hex(),
		Username:  sql.NullString{String: profile.Username, Valid: true},
		Metadata:  metadataRaw,
		Signature: "0x" + string(profile.Signature),
	}
	actual, err := testQueries.UpsertProfile(context.Background(), expected)
	require.NoError(t, err)
	require.NotEmpty(t, actual)

	require.Equal(t, expected.Address, actual.Address)
	require.Equal(t, expected.Username.String, actual.Username.String)
	require.JSONEq(t, string(expected.Metadata.RawMessage), string(actual.Metadata.RawMessage))
	require.Equal(t, expected.Signature, actual.Signature)

	return actual
}

func TestQueries_GetProfile(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	t.Run("Test get profile", func(t *testing.T) {
		expected := createRandomProfile(t)
		actual, error := testQueries.GetProfile(context.Background(), expected.Address)

		require.NoError(t, error)
		require.NotEmpty(t, actual)

		require.Equal(t, expected.Address, actual.Address)
		require.Equal(t, expected.Username.String, actual.Username.String)
		require.JSONEq(t, string(expected.Metadata.RawMessage), string(actual.Metadata.RawMessage))
		require.Equal(t, expected.Signature, actual.Signature)
	})
}

func TestQueries_UpsertProfile(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	t.Run("Test create new profile", func(t *testing.T) {
		createRandomProfile(t)
	})

	t.Run("Test update existed profile", func(t *testing.T) {
		profile := createRandomProfile(t)
		json, _ := json.Marshal(profile.Metadata)
		metadataRaw := pqtype.NullRawMessage{
			RawMessage: json,
			Valid:      true,
		}
		expected := UpsertProfileParams{
			Address:   profile.Address,
			Username:  sql.NullString{String: "TRUONG PHU HUNG", Valid: true},
			Metadata:  metadataRaw,
			Signature: "0x" + string(profile.Signature),
		}
		actual, err := testQueries.UpsertProfile(context.Background(), expected)
		require.NoError(t, err)
		require.NotEmpty(t, profile)

		require.Equal(t, expected.Address, actual.Address)
		require.Equal(t, expected.Username.String, actual.Username.String)
		require.JSONEq(t, string(expected.Metadata.RawMessage), string(actual.Metadata.RawMessage))
		require.Equal(t, expected.Signature, actual.Signature)
	})
}

func RandomHex(n int) string {
	const letterBytes = "abcdef0123456789"
	var sb strings.Builder
	for i := 0; i < n; i++ {
		c := letterBytes[rand.Intn(len(letterBytes))]
		sb.WriteByte(c)
	}
	return "0x" + sb.String()
}

func RandomAddress() string {
	return RandomHex(40)
}

func RandomString(n int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len(alphabet))]
		sb.WriteByte(c)
	}
	return sb.String()
}
