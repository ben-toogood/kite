package handler

import (
	"context"

	"gorm.io/gorm"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/likes"
	"github.com/ben-toogood/kite/likes/model"
)

// Unlike a resource
func (l *Likes) Unlike(ctx context.Context, req *likes.UnlikeRequest) (*likes.UnlikeResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// construct the like and delete it from the db
	like := &model.Like{
		ResourceType: req.ResourceType,
		ResourceID:   req.ResourceId,
		UserID:       req.UserId,
	}
	if err := l.DB.Where(like).Delete(&model.Like{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, database.TranslateError(err)
	}

	return &likes.UnlikeResponse{}, nil
}
