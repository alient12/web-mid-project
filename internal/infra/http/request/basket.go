package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type BasketCreate struct {
	Data  string `json:"data,omitempty,max=512" validate:"required"`
	State bool   `json:"state,omitempty" validate:"required"`
}

func (bc BasketCreate) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(bc); err != nil {
		return fmt.Errorf("create request validation failed %w", err)
	}

	return nil
}
