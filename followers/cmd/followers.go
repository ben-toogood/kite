package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/followers"
	"github.com/ben-toogood/kite/followers/handler"
	"github.com/ben-toogood/kite/followers/model"
	"google.golang.org/grpc"
)

func main() {
	// connect to the database
	db, err := database.GetDB(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	if err := db.AutoMigrate(&model.Follow{}); err != nil {
		fmt.Println(err)
	}

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
	followers.RegisterFollowersServiceServer(grpcServer, &handler.Followers{DB: db})
	fmt.Printf("Starting server on :%v\n", port)
	grpcServer.Serve(lis)
}
