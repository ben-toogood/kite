package posts

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . PostsServiceClient

import (
	"os"
	sync "sync"

	grpc "google.golang.org/grpc"
)

const (
	defaultAddress = "posts:8080"
)

var (
	once sync.Once
	cli  PostsServiceClient
)

func NewClient() PostsServiceClient {
	once.Do(func() {
		addr := defaultAddress
		if a := os.Getenv("POSTS_ADDRESS"); len(a) > 0 {
			addr = a
		}
		conn, _ := grpc.Dial(addr, grpc.WithInsecure())
		cli = NewPostsServiceClient(conn)
	})
	return cli
}
