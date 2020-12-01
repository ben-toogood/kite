package followers

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (a *GetFollowersRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.ResourceId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.ResourceType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
	)
}

func (a *GetFollowingRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.ResourceId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.ResourceType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
	)
}

func (a *FollowRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.FollowerId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.FollowerType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
		validation.Field(&a.FollowingId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.FollowingType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
	)
}

func (a *UnfollowRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.FollowerId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.FollowerType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
		validation.Field(&a.FollowingId, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.FollowingType, validation.Required, validation.NotIn(ResourceType_RESOURCE_TYPE_UNSPECIFIED)),
	)
}
