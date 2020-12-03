package users

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (a *CreateRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.FirstName, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.LastName, validation.Required, validation.Length(1, 255)),
	)
}

func (a *GetByEmailRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Email, validation.Required, is.Email),
	)
}
