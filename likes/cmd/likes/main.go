package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/likes"
	"github.com/ben-toogood/kite/likes/handler"
	"github.com/ben-toogood/kite/likes/model"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
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
	if err := db.AutoMigrate(&model.Like{}); err != nil {
		fmt.Println(err)
	}

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
	likes.RegisterLikesServiceServer(grpcServer, &handler.Likes{
		DB: db,
	})
	fmt.Printf("Starting server on :%v\n", port)
	grpcServer.Serve(lis)
}
