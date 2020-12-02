package likes

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (a *CountRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.ResourceIds, validation.Required),
		validation.Field(&a.ResourceType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
	)
}

func (a *GetRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.ResourceIds, validation.Required),
		validation.Field(&a.ResourceType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
	)
}

func (a *LikeRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.UserId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.ResourceId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.ResourceType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
	)
}

func (a *UnlikeRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.UserId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.ResourceId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.ResourceType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
	)
}
