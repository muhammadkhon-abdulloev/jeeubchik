package utils

import (
	"context"
	"github.com/go-playground/validator/v10"
	"regexp"
)

// Use a single instance of Validate, it caches struct info
var validate *validator.Validate

const (
	CountryCode = "+992"
)

func init() {
	validate = validator.New()
}

func ValidateStruct(ctx context.Context, s interface{}) error {
	return validate.StructCtx(ctx, s)
}

func IsPhoneValid(phone string) bool {
	ok := checkCountryCode(phone)
	if !ok {
		return false
	}

	return len(phone) == 13
}

func checkCountryCode(phone string) bool {
	return phone[:4] == CountryCode
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
