package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/auth/handler"
	"github.com/ben-toogood/kite/auth/model"
	"github.com/ben-toogood/kite/common/database"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"github.com/lileio/pubsub/v2/providers/google"
	"google.golang.org/grpc"
)

func main() {
	// connect to the database
	db, err := database.GetDB(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	if err := db.AutoMigrate(&model.Token{}); err != nil {
		fmt.Println(err)
	}

	// connect to pub sub
	var ps pubsub.Provider
	if gpid := os.Getenv("GOOGLE_PUBSUB_PROJECT_ID"); len(gpid) > 0 {
		ps, err = google.NewGoogleCloud(gpid)
		if err != nil {
			fmt.Println(err)
		}
	}
	psc := &pubsub.Client{
		ServiceName: "auth",
		Provider:    ps,
		Middleware:  defaults.Middleware,
	}

	// start the server
	flag.Parse()
	port := "8080"
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	auth.RegisterAuthServiceServer(grpcServer, &handler.Auth{DB: db, PubSub: psc})
	fmt.Printf("Starting server on :%v\n", port)
	grpcServer.Serve(lis)
}
