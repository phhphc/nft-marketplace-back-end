package configs

type Config struct {
	Env              string  `mapstructure:"ENV" validate:"required,oneof=Dev Production"`
	Port             int     `mapstructure:"PORT"`
	Host             string  `mapstructure:"HOST"`
	PostgreUri       string  `mapstructure:"POSTGRE_URI" validate:"required,uri"`
	ChainUrl         string  `mapstructure:"CHAIN_URL" validate:"required,url"`
	MarketplaceAddr  string  `mapstructure:"MKP_ADDR" validate:"required,eth_addr"`
	RedisUrl         string  `mapstructure:"REDIS_URL" validate:"required"`
	RedisPass        string  `mapstructure:"REDIS_PASS" validate:"required"`
	JwtSecret        string  `mapstructure:"JWT_SECRET" validate:"required"`
	JwtExpired       string  `mapstructure:"JWT_EXPIRED" validate:"required"`
	MarketplaceAdmin string  `mapstructure:"MKP_ADMIN" validate:"required,eth_addr"`
	Royalty          float64 `mapstructure:"ROYALTY" validate:"required,gte=0,lte=1"`
}
