package dto

import "github.com/go-playground/validator/v10"

type Number struct {
	Value int64 `json:"value" validate:"required"`
}

func (number Number) Validate() error {
	validate := validator.New()
	if err := validate.Struct(number); err != nil {
		return err
	}

	return nil
}
