package handler

import (
	"context"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/followers"
	"github.com/ben-toogood/kite/followers/model"
)

// GetFollowing for a resource
func (f *Followers) GetFollowing(ctx context.Context, req *followers.GetFollowingRequest) (*followers.GetFollowingResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// construct and execute the query
	q := &model.Follow{
		FollowerType: req.ResourceType,
		FollowerID:   req.ResourceId,
	}
	var res []model.Follow
	if err := f.DB.Where(q).Find(&res).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	// serialize the result
	rsp := &followers.GetFollowingResponse{
		Following: make([]*followers.Resource, len(res)),
	}
	for i, f := range res {
		rsp.Following[i] = &followers.Resource{
			Type: f.FollowingType,
			Id:   f.FollowingID,
		}
	}

	return rsp, nil
}
