package comments

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (a *CreateRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.AuthorId, validation.Required),
		validation.Field(&a.Message, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.ResourceId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.ResourceType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
	)
}

func (a *GetRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.ResourceIds, validation.Required),
		validation.Field(&a.ResourceType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
	)
}

func (a *DeleteRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Id, validation.Required),
	)
}
