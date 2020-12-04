package server

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/posts"
	"github.com/ben-toogood/kite/posts/model"
	"github.com/lileio/lile/v2/protocopy"
	"github.com/segmentio/ksuid"
)

// Create a post
func (p *Posts) Create(ctx context.Context, req *posts.CreateRequest) (*posts.CreateResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	imgID := ksuid.New().String()
	w := p.Bucket.Object(imgID).NewWriter(ctx)
	if _, err := w.Write(req.Image); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	// construct the post and write it to the database
	post := &model.Post{
		AuthorID:    req.AuthorId,
		Description: req.Description,
		ImageID:     imgID,
	}
	if err := p.DB.Create(post).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	// get the URL for the post
	url, err := storage.SignedURL(os.Getenv("BUCKET_NAME"), imgID, &storage.SignedURLOptions{})
	if err != nil {
		return nil, err
	}

	// serialize the result
	rsp := &posts.CreateResponse{Post: &posts.Post{}}
	if err := protocopy.ToProto(post, rsp.Post); err != nil {
		return nil, err
	}
	rsp.Post.ImageUrl = url
	return rsp, nil
}
