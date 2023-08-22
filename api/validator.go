package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/urunsiyabend/simple_bank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}

var validPassword validator.Func = func(fl validator.FieldLevel) bool {
	if password, ok := fl.Field().Interface().(string); ok {
		return util.IsStrongPassword(password)
	}
	return false
}
