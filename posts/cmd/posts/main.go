package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"cloud.google.com/go/storage"
	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/posts"
	"github.com/ben-toogood/kite/posts/model"
	"github.com/ben-toogood/kite/posts/server"
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
	if err := db.AutoMigrate(&model.Post{}); err != nil {
		fmt.Println(err)
	}

	// connect to google cloud storage
	client, err := storage.NewClient(context.Background())
	if err != nil {
		panic(err)
	}
	bucket := client.Bucket(os.Getenv("BUCKET_NAME"))

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
	posts.RegisterPostsServiceServer(grpcServer, &server.Posts{
		Bucket: bucket,
		DB:     db,
	})
	fmt.Printf("Starting server on :%v\n", port)
	grpcServer.Serve(lis)
}
