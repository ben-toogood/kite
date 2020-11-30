package handler

import (
	"context"

	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/comments/model"
	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
)

// Delete a comment
func (c *Comments) Delete(ctx context.Context, req *comments.DeleteRequest) (*comments.DeleteResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// delete the comment
	if err := c.DB.Delete(&model.Comment{ID: req.Id}).Error; err != nil {
		return nil, database.TranslateErrors(err)
	}

	return &comments.DeleteResponse{}, nil
}
