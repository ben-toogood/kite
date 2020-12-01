package resolvers

import (
	"context"

	"github.com/ben-toogood/kite/users"
	"github.com/graph-gophers/graphql-go"
)

type User struct {
	u *users.User
}

func (u *User) ID() graphql.ID {
	return graphql.ID(u.u.Id)
}

func (u *User) FirstName() string {
	return u.u.FirstName
}

func (u *User) LastName() string {
	return u.u.LastName
}

func (r *Resolver) User(ctx context.Context, args struct{ ID graphql.ID }) (*User, error) {
	rsp, err := r.Users.Get(ctx, &users.GetRequest{
		Ids: []string{string(args.ID)},
	})
	if err != nil {
		return nil, err
	}

	return &User{u: rsp.Users[string(args.ID)]}, nil
}
