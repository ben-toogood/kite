package graph

import (
	"context"

	"github.com/ben-toogood/kite/auth"
)

func (r *mutationResolver) RefreshTokens(ctx context.Context, refreshToken string) (*Tokens, error) {
	res, err := r.Auth.Refresh(ctx, &auth.RefreshRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  res.Token.AccessToken,
		RefreshToken: res.Token.RefreshToken,
	}, nil
}
