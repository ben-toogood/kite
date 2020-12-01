package resolvers

import (
	"context"

	"github.com/ben-toogood/kite/users"
)

type SignupInput struct {
	FirstName string
	LastName  string
	Email     string
}

func (r *Resolver) Signup(ctx context.Context, input SignupInput) (*User, error) {
	rsp, err := users.NewClient().Create(ctx, &users.CreateRequest{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	})
	if err != nil {
		return nil, err
	}

	return &User{u: rsp.User}, nil
}
