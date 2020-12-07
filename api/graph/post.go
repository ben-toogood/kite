package graph

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ben-toogood/kite/followers"
	"github.com/ben-toogood/kite/posts"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Post struct {
	p *posts.Post
	// AuthorId string `json:"author"`
}

func (p *Post) ID() string {
	return p.p.Id
}

func (p *Post) ImageURL() string {
	return p.p.ImageUrl
}

func (p *Post) Description() string {
	return p.p.Description
}

func (r *mutationResolver) CreatePost(ctx context.Context, description string, image graphql.Upload) (*Post, error) {
	userID := UserIDFromContext(ctx)
	if len(userID) == 0 {
		return nil, errUnauthorized
	}

	f, err := ioutil.ReadAll(image.File)
	if err != nil {
		return nil, err
	}

	rsp, err := r.Posts.Create(ctx, &posts.CreateRequest{
		AuthorId:    userID,
		Description: description,
		Image:       f,
	})
	if err != nil {
		return nil, err
	}

	return &Post{p: rsp.Post}, nil
}

func (r *queryResolver) GetPosts(ctx context.Context, createdBefore *string) ([]*Post, error) {
	userID := UserIDFromContext(ctx)
	if len(userID) == 0 {
		return nil, errUnauthorized
	}

	// get the users they're following
	fRsp, err := r.Followers.GetFollowing(ctx, &followers.GetFollowingRequest{
		ResourceId: userID, ResourceType: followers.ResourceType_RESOURCE_TYPE_USER,
	})
	if err != nil {
		return nil, err
	}
	authorIDs := make([]string, len(fRsp.Following))
	for i, r := range fRsp.Following {
		authorIDs[i] = r.Id
	}
	authorIDs = append(authorIDs, userID)

	// lookup the posts authored by these users
	pr := &posts.ListRequest{AuthorIds: authorIDs}
	if createdBefore != nil {
		t, err := time.Parse(time.RFC3339, *createdBefore)
		if err != nil {
			return nil, err
		}

		pr.CreatedBefore = timestamppb.New(t)
	}
	pRsp, err := r.Posts.List(ctx, pr)
	if err != nil {
		return nil, err
	}

	// construst the response
	posts := make([]*Post, len(pRsp.Posts))
	for i, p := range pRsp.Posts {
		posts[i] = &Post{p: p}
	}
	return posts, nil
}

func (r *postResolver) Author(ctx context.Context, obj *Post) (*User, error) {
	return LoadersFor(ctx).UserById.Load(obj.p.AuthorId)
}
