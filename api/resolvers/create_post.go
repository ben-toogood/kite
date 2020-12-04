package resolvers

import (
	"context"
	"errors"
	"io/ioutil"

	"github.com/ben-toogood/kite/posts"
)

type CreatePostInput struct {
	Image       *GraphQLUpload
	Description string
}

var errUnauthorized = errors.New("Unauthorized")

func (r *Resolver) CreatePost(ctx context.Context, input CreatePostInput) (*Post, error) {
	userID := UserIDFromContext(ctx)
	if len(userID) == 0 {
		return nil, errUnauthorized
	}

	bytes, err := ioutil.ReadFile(input.Image.Filepath)
	if err != nil {
		return nil, err
	}

	rsp, err := r.Posts.Create(ctx, &posts.CreateRequest{
		AuthorId:    userID,
		Description: input.Description,
		Image:       bytes,
	})
	if err != nil {
		return nil, err
	}

	return &Post{p: rsp.Post}, nil
}
