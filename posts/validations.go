package posts

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (a *CreateRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.AuthorId, validation.Required),
		validation.Field(&a.Description, validation.Required),
		validation.Field(&a.Image, validation.Required),
	)
}

func (a *ListRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.AuthorIds, validation.Required),
	)
}
