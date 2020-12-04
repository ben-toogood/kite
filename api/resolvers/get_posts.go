package resolvers

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ben-toogood/kite/followers"
	"github.com/ben-toogood/kite/posts"
)

type GetPostsInput struct {
	CreatedBefore *Timestamp
}

func (r *Resolver) GetPosts(ctx context.Context, input GetPostsInput) (*[]*Post, error) {
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
	if input.CreatedBefore != nil {
		pr.CreatedBefore = timestamppb.New(input.CreatedBefore.Time)
	}
	pRsp, err := r.Posts.List(ctx, pr)
	if err != nil {
		return nil, err
	}

	// construst the response
	posts := make([]*Post, len(pRsp.Posts))
	for i, p := range pRsp.Posts {
		posts[i] = &Post{p}
	}
	return &posts, nil
}
