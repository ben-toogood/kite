package graph

import (
	"context"
	"time"

	"github.com/ben-toogood/kite/users"
)

func NewUserLoaderWithCtx(r *Resolver, ctx context.Context) *UserLoader {
	return NewUserLoader(
		UserLoaderConfig{
			Wait:     1 * time.Millisecond,
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

func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	return LoadersFor(ctx).UserById.Load(id)
}
