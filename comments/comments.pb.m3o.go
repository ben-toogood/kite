package comments

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . CommentsServiceClient

import (
	sync "sync"

	grpc "google.golang.org/grpc"
)

var (
	once sync.Once
	cli  CommentsServiceClient
)

func NewClient() CommentsServiceClient {
	once.Do(func() {
		conn, _ := grpc.Dial("Comments", grpc.WithInsecure())
		cli = NewCommentsServiceClient(conn)
	})
	return cli
}
