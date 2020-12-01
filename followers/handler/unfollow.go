package handler

import (
	"context"

	"gorm.io/gorm"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/followers"
	"github.com/ben-toogood/kite/followers/model"
)

// Unfollow a resource
func (f *Followers) Unfollow(ctx context.Context, req *followers.UnfollowRequest) (*followers.UnfollowResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// construct the follow object and delete it from the database
	fo := &model.Follow{
		FollowerType:  req.FollowerType,
		FollowerID:    req.FollowerId,
		FollowingType: req.FollowingType,
		FollowingID:   req.FollowingId,
	}
	if err := f.DB.Where(fo).Delete(&model.Follow{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, database.TranslateError(err)
	}

	return &followers.UnfollowResponse{}, nil
}
