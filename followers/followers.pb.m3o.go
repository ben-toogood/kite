package followers

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . UsersServiceClient

import (
	"os"
	sync "sync"

	grpc "google.golang.org/grpc"
)

const (
	defaultAddress = "followers:8080"
)

var (
	once sync.Once
	cli  FollowersServiceClient
)

func NewClient() FollowersServiceClient {
	once.Do(func() {
		addr := defaultAddress
		if a := os.Getenv("FOLLOWERS_ADDRESS"); len(a) > 0 {
			addr = a
		}
		conn, _ := grpc.Dial(addr, grpc.WithInsecure())
		cli = NewFollowersServiceClient(conn)
	})
	return cli
}
