package handler

import (
	"context"

	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/comments/model"
	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/lileio/lile/v2/protocopy"
)

// Get comments for a resource
func (c *Comments) Get(ctx context.Context, req *comments.GetRequest) (*comments.GetResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// query the database
	var cmts []model.Comment
	query := c.DB.Where("resource_type = ? AND resource_id IN (?)", req.ResourceType, req.ResourceIds)
	if err := query.Find(&cmts).Error; err != nil {
		return nil, database.TranslateErrors(err)
	}

	// group the results by resource_id
	rsp := &comments.GetResponse{
		Resources: make(map[string]*comments.Resource),
	}
	for _, c := range cmts {
		var cmt comments.Comment
		if err := protocopy.ToProto(c, &cmt); err != nil {
			return nil, err
		}

		if cs, ok := rsp.Resources[c.ResourceID]; ok {
			cs.Comments = append(cs.Comments, &cmt)
		} else {
			rsp.Resources[c.ResourceID] = &comments.Resource{
				Comments: []*comments.Comment{&cmt},
			}
		}
	}

	return rsp, nil
}
