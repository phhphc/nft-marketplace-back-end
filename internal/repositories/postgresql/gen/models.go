// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package gen

import (
	"database/sql"

	"github.com/tabbed/pqtype"
)

type Category struct {
	ID   int32
	Name string
}

type Collection struct {
	Token         string
	Owner         string
	Name          string
	Description   string
	Metadata      pqtype.NullRawMessage
	Category      int32
	CreatedAt     sql.NullTime
	LastSyncBlock int64
}

type ConsiderationItem struct {
	ID          int64
	OrderHash   string
	ItemType    int32
	Token       string
	Identifier  string
	Amount      sql.NullString
	StartAmount sql.NullString
	EndAmount   sql.NullString
	Recipient   string
}

type Event struct {
	ID        int32
	Name      string
	Token     string
	TokenID   string
	Quantity  sql.NullInt32
	Type      sql.NullString
	Price     sql.NullString
	From      string
	To        sql.NullString
	Date      sql.NullTime
	TxHash    sql.NullString
	OrderHash sql.NullString
}

type Marketplace struct {
	LastSyncBlock int64
}

type MarketplaceSetting struct {
	ID          int32
	Marketplace string
	Beneficiary string
	Royalty     string
}

type Nft struct {
	Token       string
	Identifier  string
	Owner       string
	TokenUri    sql.NullString
	Metadata    pqtype.NullRawMessage
	IsBurned    bool
	IsHidden    bool
	BlockNumber string
	TxIndex     int64
}

type Notification struct {
	ID        int32
	Info      string
	EventName string
	OrderHash string
	Address   string
	IsViewed  sql.NullBool
}

type OfferItem struct {
	ID          int64
	OrderHash   string
	ItemType    int32
	Token       string
	Identifier  string
	Amount      sql.NullString
	StartAmount sql.NullString
	EndAmount   sql.NullString
}

type Order struct {
	OrderHash   string
	Offerer     string
	Recipient   sql.NullString
	Salt        sql.NullString
	StartTime   sql.NullString
	EndTime     sql.NullString
	Signature   sql.NullString
	IsCancelled bool
	IsValidated bool
	IsFulfilled bool
	IsInvalid   bool
}
