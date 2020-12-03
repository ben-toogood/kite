package resolvers

import (
	"context"
	"time"

	"github.com/ben-toogood/kite/users"
	"github.com/graph-gophers/graphql-go"
)

func NewUserLoaderWithCtx(r *Resolver, ctx context.Context) *UserLoader {
	return NewUserLoader(
		UserLoaderConfig{
			Wait:     2 * time.Millisecond,
			MaxBatch: 100,
			Fetch: func(keys []string) ([]*User, []error) {
				us := make([]*User, len(keys))
				errors := make([]error, len(keys))

				rsp, err := r.Users.Get(ctx, &users.GetRequest{
					Ids: keys,
				})
				if err != nil {
					return nil, []error{err}
				}

				for i, key := range keys {
					us[i] = &User{u: rsp.Users[key]}
				}

				return us, errors
			},
		})
}

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
	return LoadersFor(ctx).UserById.Load(string(args.ID))
}
