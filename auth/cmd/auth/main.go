package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/auth/handler"
	"github.com/ben-toogood/kite/auth/model"
	"github.com/ben-toogood/kite/common/database"
	"github.com/form3tech-oss/jwt-go"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"github.com/lileio/pubsub/v2/providers/google"
	"github.com/sendgrid/sendgrid-go"
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

	// Load the certs
	keyPath := os.Getenv("PRIVATE_KEY_FILEPATH")
	if len(keyPath) == 0 {
		panic("Missing PRIVATE_KEY_FILEPATH")
	}
	file, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(file)
	if err != nil {
		panic(err)
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
	auth.RegisterAuthServiceServer(grpcServer, &handler.Auth{
		DB:         db,
		PubSub:     psc,
		PrivateKey: key,
		Sendgrid:   sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY")),
	})
	fmt.Printf("Starting server on :%v\n", port)
	grpcServer.Serve(lis)
}
