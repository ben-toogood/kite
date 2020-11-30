package users

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fake/users.go . UsersServiceClient

import (
	sync "sync"

	grpc "google.golang.org/grpc"
)

var (
	once sync.Once
	cli  UsersServiceClient
)

func NewClient() UsersServiceClient {
	once.Do(func() {
		conn, _ := grpc.Dial("users", grpc.WithInsecure())
		cli = NewUsersServiceClient(conn)
	})
	return cli
}
