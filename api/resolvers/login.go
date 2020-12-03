package resolvers

import (
	"context"

	"github.com/ben-toogood/kite/auth"
)

type LoginInput struct {
	Email string
}

func (r *Resolver) Login(ctx context.Context, input LoginInput) (*bool, error) {
	_, err := r.Auth.Login(ctx, &auth.LoginRequest{
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}

	rs := true
	return &rs, nil
}
