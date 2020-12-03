package comments

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . CommentsServiceClient

import (
	"os"
	sync "sync"

	grpc "google.golang.org/grpc"
)

const (
	defaultAddress = "comments:8080"
)

var (
	once sync.Once
	cli  CommentsServiceClient
)

func NewClient() CommentsServiceClient {
	once.Do(func() {
		addr := defaultAddress
		if a := os.Getenv("COMMENTS_ADDRESS"); len(a) > 0 {
			addr = a
		}
		conn, _ := grpc.Dial(addr, grpc.WithInsecure())
		cli = NewCommentsServiceClient(conn)
	})
	return cli
}
