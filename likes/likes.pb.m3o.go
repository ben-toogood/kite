package likes

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . LikesServiceClient

import (
	"os"
	sync "sync"

	grpc "google.golang.org/grpc"
)

const (
	defaultAddress = "likes:8080"
)

var (
	once sync.Once
	cli  LikesServiceClient
)

func NewClient() LikesServiceClient {
	once.Do(func() {
		addr := defaultAddress
		if a := os.Getenv("LIKES_ADDRESS"); len(a) > 0 {
			addr = a
		}
		conn, _ := grpc.Dial(addr, grpc.WithInsecure())
		cli = NewLikesServiceClient(conn)
	})
	return cli
}
