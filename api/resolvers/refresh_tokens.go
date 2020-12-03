package resolvers

import (
	"context"

	"github.com/ben-toogood/kite/auth"
)

type RefreshTokensInput struct {
	RefreshToken string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func (r *Resolver) RefreshTokens(ctx context.Context, input RefreshTokensInput) (*Tokens, error) {
	res, err := r.Auth.Refresh(ctx, &auth.RefreshRequest{
		RefreshToken: input.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  res.Token.AccessToken,
		RefreshToken: res.Token.RefreshToken,
	}, nil
}
