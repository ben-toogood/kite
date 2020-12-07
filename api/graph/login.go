package graph

import (
	"context"

	"github.com/ben-toogood/kite/auth"
)

func (r *mutationResolver) Login(ctx context.Context, email string) (*bool, error) {
	_, err := r.Auth.Login(ctx, &auth.LoginRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	rs := true
	return &rs, nil
}
