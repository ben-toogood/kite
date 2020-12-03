package auth

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (a *LoginRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Email, validation.Required, is.Email),
	)
}

func (a *RefreshRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.RefreshToken, validation.Required),
	)
}

func (a *RevokeRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.UserId, validation.Required, validation.Length(1, 255)),
	)
}
