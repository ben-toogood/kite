package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/model"
	"github.com/ben-toogood/kite/users/server"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"github.com/lileio/pubsub/v2/providers/google"
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"google.golang.org/grpc"
)

func main() {
	// Jeager
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		// parsing errors might happen here, such as when we get a string where we expect a number
		log.Printf("Could not parse Jaeger env vars: %s", err.Error())
		return
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaegerlog.StdLogger))
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	// connect to the database
	db, err := database.GetDB(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		fmt.Println(err)
	}

	// connect to pub sub
	var ps pubsub.Provider
	if gpid := os.Getenv("GOOGLE_PUBSUB_PROJECT_ID"); len(gpid) > 0 {
		ps, err = google.NewGoogleCloud(gpid)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		panic("no GOOGLE_PUBSUB_PROJECT_ID set! 😂")
	}

	// Pubsub
	pubsub.SetClient(&pubsub.Client{
		ServiceName: "users",
		Provider:    ps,
		Middleware:  defaults.Middleware,
	})

	// start the server
	flag.Parse()
	port := "8080"
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())),
	)
	grpcServer := grpc.NewServer(opts...)
	users.RegisterUsersServiceServer(grpcServer, &server.Users{})
	fmt.Printf("Starting server on :%v\n", port)
	grpcServer.Serve(lis)
}
