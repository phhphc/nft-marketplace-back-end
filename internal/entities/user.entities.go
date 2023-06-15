package entities

import "github.com/golang-jwt/jwt"

type User struct {
	Address string `json:"address"`
	Nonce   string `json:"nonce"`
	Roles   []Role `json:"roles"`
	IsBlock bool   `json:"is_block"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Claims struct {
	Address string   `json:"address"`
	Nonce   string   `json:"nonce"`
	Roles   []string `json:"roles"`
	jwt.StandardClaims
}
