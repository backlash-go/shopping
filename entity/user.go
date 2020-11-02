package entity

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

type LoginRequestParam struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}
