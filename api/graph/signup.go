package graph

import (
	"context"

	"github.com/ben-toogood/kite/users"
)

func (r *mutationResolver) Signup(ctx context.Context, firstName string, lastName string, email string) (*User, error) {
	rsp, err := r.Users.Create(ctx, &users.CreateRequest{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	})
	if err != nil {
		return nil, err
	}

	return &User{u: rsp.User}, nil
}
