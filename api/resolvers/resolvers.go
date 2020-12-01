package resolvers

import (
	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/users"
)

type Resolver struct {
	Users    users.UsersServiceClient
	Comments comments.CommentsServiceClient
}
