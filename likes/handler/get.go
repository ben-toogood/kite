package handler

import (
	"context"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/likes"
	"github.com/ben-toogood/kite/likes/model"
	"github.com/lileio/lile/v2/protocopy"
)

// Get the the users who liked a list of resources
func (l *Likes) Get(ctx context.Context, req *likes.GetRequest) (*likes.GetResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// execute the query
	var data []model.Like
	q := l.DB.Where("resource_type = ? AND resource_id IN (?)", req.ResourceType, req.ResourceIds)
	if err := q.Find(&data).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	// serialize the result
	rsp := &likes.GetResponse{
		Resources: make(map[string]*likes.ResourceLikes),
	}
	for _, l := range data {
		var ls likes.Like
		if err := protocopy.ToProto(l, &ls); err != nil {
			return nil, err
		}
		if r, ok := rsp.Resources[l.ResourceID]; ok {
			rsp.Resources[l.ResourceID].Likes = append(r.Likes, &ls)
		} else {
			rsp.Resources[l.ResourceID] = &likes.ResourceLikes{
				Likes: []*likes.Like{&ls},
			}
		}
	}
	return rsp, nil
}
