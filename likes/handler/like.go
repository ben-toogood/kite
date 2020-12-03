package handler

import (
	"context"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/likes"
	"github.com/ben-toogood/kite/likes/model"
)

// Like a resource
func (l *Likes) Like(ctx context.Context, req *likes.LikeRequest) (*likes.LikeResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// construct the like and write to the db
	like := &model.Like{
		ResourceType: req.ResourceType,
		ResourceID:   req.ResourceId,
		UserID:       req.UserId,
	}
	if err := l.DB.Create(like).Error; err != nil {
		if verr := database.TranslateError(err); verr != database.ErrDuplicate {
			return nil, verr
		}
	}

	return &likes.LikeResponse{}, nil
}
