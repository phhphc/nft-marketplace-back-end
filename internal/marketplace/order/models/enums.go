package models

type EnumItemType int

const (
	NAIVE EnumItemType = iota
	ERC20
	ERC721
	ERC1155
)

type EnumOrderType int

const (
	FULL_OPEN EnumOrderType = iota
	FULL_RESTRICTED
	PARTIAL_RESTRICTED
	CONTRACT
)
