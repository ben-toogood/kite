package graph

import (
	"crypto/rsa"
	"errors"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/followers"
	"github.com/ben-toogood/kite/posts"
	"github.com/ben-toogood/kite/users"
)

//go:generate go run github.com/vektah/dataloaden UserLoader string "*github.com/ben-toogood/kite/api/graph.User"

var errUnauthorized = errors.New("Unauthorized")

type Resolver struct {
	Auth      auth.AuthServiceClient
	Comments  comments.CommentsServiceClient
	Followers followers.FollowersServiceClient
	Users     users.UsersServiceClient
	Posts     posts.PostsServiceClient
	PublicKey *rsa.PublicKey
}
