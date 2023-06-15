package httpApi

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

var validate *validator.Validate

func NewValidator() *Validator {
	validate = validator.New()
	validate.RegisterValidation("eth_hash", func(fl validator.FieldLevel) bool {
		const ethHashRegexString = `^0x[0-9a-fA-F]{64}$`
		var ethHashRegex = regexp.MustCompile(ethHashRegexString)

		return ethHashRegex.MatchString(fl.Field().String())
	})

	return &Validator{
		validator: validate,
	}
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
