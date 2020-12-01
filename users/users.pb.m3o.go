package users

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . UsersServiceClient

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
