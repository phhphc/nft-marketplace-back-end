package entities

type User struct {
	Address string   `json:"address"`
	Nonce   string   `json:"nonce"`
	Roles   []string `json:"roles"`
}
