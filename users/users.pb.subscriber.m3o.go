package users

import (
	"context"

	"github.com/lileio/pubsub/v2"
)

func SubscribeUsersCreated(c *pubsub.Client, name string, f func(ctx context.Context, obj *User, psMsg *pubsub.Msg) error) {
	// This should allow you to set this stuff
	c.On(pubsub.HandlerOptions{
		Topic:   UserCreated,
		Name:    name,
		Handler: f,
		AutoAck: true,
	})
}
