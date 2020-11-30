package graph

import (
	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/fake"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Users users.UsersServiceClient
}

func testResolver() *Resolver {
	return &Resolver{
		Users: new(fake.FakeUsersServiceClient),
	}
}
