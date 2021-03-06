package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/comments/handler"
	"github.com/ben-toogood/kite/comments/model"
	"github.com/ben-toogood/kite/common/database"
	"google.golang.org/grpc"
)

func main() {
	// connect to the database
	db, err := database.GetDB(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	if err := db.AutoMigrate(&model.Comment{}); err != nil {
		fmt.Println(err)
	}

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
	grpcServer := grpc.NewServer(opts...)
	comments.RegisterCommentsServiceServer(grpcServer, &handler.Comments{DB: db})
	fmt.Printf("Starting server on :%v\n", port)
	grpcServer.Serve(lis)
}
