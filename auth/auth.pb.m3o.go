package auth

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . AuthServiceClient

import (
	"os"
	sync "sync"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"

	grpc "google.golang.org/grpc"
)

const (
	defaultAddress = "auth:8080"
)

var (
	once sync.Once
	cli  AuthServiceClient
)

func NewClient() AuthServiceClient {
	once.Do(func() {
		addr := defaultAddress
		if a := os.Getenv("AUTH_ADDRESS"); len(a) > 0 {
			addr = a
		}
		conn, _ := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
		cli = NewAuthServiceClient(conn)
	})
	return cli
}
