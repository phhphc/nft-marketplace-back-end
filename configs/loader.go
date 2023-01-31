package configs

import (
	_ "embed"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

//go:embed default.yaml
var defaultConfig string

func LoadConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(strings.NewReader(defaultConfig)); err != nil {
		return nil, err
	}
	viper.AutomaticEnv()

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
