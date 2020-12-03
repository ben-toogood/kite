package resolvers

import (
	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/users"
)

type Resolver struct {
	Users    users.UsersServiceClient
	Auth     auth.AuthServiceClient
	Comments comments.CommentsServiceClient
}
