package server

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/posts"
	"github.com/ben-toogood/kite/posts/model"
	"github.com/lileio/lile/v2/protocopy"
)

const postsLimit = 25

// List posts
func (p *Posts) List(ctx context.Context, req *posts.ListRequest) (*posts.ListResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// construct the query
	q := p.DB.Where("author_id IN (?)", req.AuthorIds)
	q = q.Order("created_at DESC").Limit(postsLimit)
	if req.CreatedBefore != nil {
		q = q.Where("created_at < ?", req.CreatedBefore.AsTime())
	}

	// execute the query
	var data []model.Post
	if err := q.Find(&data).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	// serialize the result
	rsp := &posts.ListResponse{
		Posts: make([]*posts.Post, len(data)),
	}
	for i, ps := range data {
		// get the URL for the post
		url, err := storage.SignedURL(os.Getenv("BUCKET_NAME"), ps.ImageID, &storage.SignedURLOptions{
			GoogleAccessID: p.GoogleAccessID,
			PrivateKey:     p.GooglePrivateKey,
			Method:         "GET",
			Expires:        time.Now().Add(1 * time.Hour),
		})
		if err != nil {
			return nil, err
		}

		var post posts.Post
		if err := protocopy.ToProto(ps, &post); err != nil {
			return nil, err
		}
		post.ImageUrl = url

		rsp.Posts[i] = &post
	}
	return rsp, nil
}
