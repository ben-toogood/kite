package handler

import (
	"context"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/followers"
	"github.com/ben-toogood/kite/followers/model"
)

// GetFollowers for a resource
func (f *Followers) GetFollowers(ctx context.Context, req *followers.GetFollowersRequest) (*followers.GetFollowersResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// construct and execute the query
	q := &model.Follow{
		FollowingType: req.ResourceType,
		FollowingID:   req.ResourceId,
	}
	var res []model.Follow
	if err := f.DB.Where(q).Find(&res).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	// serialize the result
	rsp := &followers.GetFollowersResponse{
		Followers: make([]*followers.Resource, len(res)),
	}
	for i, f := range res {
		rsp.Followers[i] = &followers.Resource{
			Type: f.FollowerType,
			Id:   f.FollowerID,
		}
	}

	return rsp, nil
}
