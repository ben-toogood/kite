package graph

import "github.com/ben-toogood/kite/users"

type User struct {
	u *users.User
}

func (u *User) ID() string {
	return u.u.Id
}

func (u *User) FirstName() string {
	return u.u.FirstName
}

func (u *User) LastName() string {
	return u.u.LastName
}
