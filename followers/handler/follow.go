package handler

import (
	"context"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/followers"
	"github.com/ben-toogood/kite/followers/model"
)

// Follow a resource
func (f *Followers) Follow(ctx context.Context, req *followers.FollowRequest) (*followers.FollowResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// construct the follow object and write it to the database
	fo := &model.Follow{
		FollowerType:  req.FollowerType,
		FollowerID:    req.FollowerId,
		FollowingType: req.FollowingType,
		FollowingID:   req.FollowingId,
	}
	if err := f.DB.Create(&fo).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	return &followers.FollowResponse{}, nil
}
