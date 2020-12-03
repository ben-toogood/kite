package users

import (
	"context"

	"github.com/lileio/pubsub/v2"
)

type SerializeableUser interface {
	Serialize() (*User, error)
}

const UserCreated = "users.user.created"

func PublishUserCreated(ctx context.Context, u SerializeableUser) error {
	usr, err := u.Serialize()
	if err != nil {
		return err
	}
	res := pubsub.Publish(ctx, UserCreated, usr)
	<-res.Ready
	return res.Err
}
