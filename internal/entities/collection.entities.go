package entities

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Collection struct {
	Token common.Address
	Owner common.Address

	Name        string
	Description string
	Metadata    map[string]any
	Category    string
	CreatedAt   time.Time

	//"metadata"    jsonb
}
