package resolvers

import (
	"context"
	"net/http"
)

const loadersKey = "loaders"

type Loaders struct {
	UserById *UserLoader
}

func ContextWithLoaders(r *Resolver, ctx context.Context) context.Context {
	return context.WithValue(ctx, loadersKey, &Loaders{
		UserById: NewUserLoaderWithCtx(r, ctx),
	})
}

func WithLoaders(res *Resolver, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(ContextWithLoaders(res, r.Context()))
		next.ServeHTTP(w, r)
	})
}

func LoadersFor(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
