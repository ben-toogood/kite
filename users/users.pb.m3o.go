package users

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . UsersServiceClient

import (
	"os"
	sync "sync"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"

	grpc "google.golang.org/grpc"
)

const (
	defaultAddress = "users:8080"
)

var (
	once sync.Once
	cli  UsersServiceClient
)

func NewClient() UsersServiceClient {
	once.Do(func() {
		addr := defaultAddress
		if a := os.Getenv("USERS_ADDRESS"); len(a) > 0 {
			addr = a
		}
		conn, _ := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
		cli = NewUsersServiceClient(conn)
	})
	return cli
}
