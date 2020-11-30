package handler

import (
	"context"

	"github.com/ben-toogood/kite/comments/model"
	pb "github.com/ben-toogood/kite/comments/proto"
	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/lileio/lile/v2/protocopy"
)

// Create a user
func (c *Comments) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// construct the object and write it to the database
	cmt := model.Comment{
		AuthorID:     req.AuthorId,
		ResourceType: req.ResourceType,
		ResourceID:   req.ResourceId,
		Message:      req.Message,
	}
	if err := c.DB.Create(&cmt).Error; err != nil {
		return nil, database.TranslateErrors(err)
	}

	// serialize the result
	rsp := pb.CreateResponse{Comment: &pb.Comment{}}
	if err := protocopy.ToProto(cmt, rsp.Comment); err != nil {
		return nil, err
	}
	return &rsp, nil
}
