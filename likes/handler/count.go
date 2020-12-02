package handler

import (
	"context"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/likes"
)

func (l *Likes) Count(ctx context.Context, req *likes.CountRequest) (*likes.CountResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// count the likes for the posts
	var resources []struct {
		ResourceID string
		Count      int32
	}
	query := l.DB.Where("resource_type = ? AND resource_id IN (?)", req.ResourceType, req.ResourceIds)
	query = query.Select("resource_id, COUNT(user_id)").Group("resource_id")
	if err := query.Scan(&resources).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	// format the result
	rsp := &likes.CountResponse{
		Counts: make(map[string]int32, len(resources)),
	}
	for _, c := range resources {
		rsp.Counts[c.ResourceID] = c.Count
	}
	return rsp, nil
}
