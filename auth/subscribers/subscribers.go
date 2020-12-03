package subscribers

import (
	"context"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/auth/handler"
	"github.com/ben-toogood/kite/users"
	"github.com/lileio/pubsub/v2"
)

type AuthServiceSubscriber struct {
	Handler *handler.Auth
}

func (s *AuthServiceSubscriber) Setup(c *pubsub.Client) {
	users.SubscribeUsersCreated(c, "send-auth-token", s.UserCreated)
}

func (s *AuthServiceSubscriber) UserCreated(ctx context.Context, u *users.User, _ *pubsub.Msg) error {
	_, err := s.Handler.Login(ctx, &auth.LoginRequest{UserId: u.Id, Email: u.Email})
	return err
}
