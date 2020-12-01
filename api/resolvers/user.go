package resolvers

import (
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
