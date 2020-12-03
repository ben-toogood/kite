package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/auth/handler"
	"github.com/ben-toogood/kite/auth/model"
	"github.com/ben-toogood/kite/auth/subscribers"
	"github.com/ben-toogood/kite/common/database"
	"github.com/form3tech-oss/jwt-go"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"github.com/lileio/pubsub/v2/providers/google"
	"github.com/opentracing/opentracing-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
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

	logrus.SetLevel(logrus.DebugLevel)

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
	pubsub.SetClient(psc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

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
	opts = append(opts, grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())),
	)

	h := &handler.Auth{
		DB:         db,
		PubSub:     psc,
		PrivateKey: key,
		Sendgrid:   sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY")),
	}

	grpcServer := grpc.NewServer(opts...)
	auth.RegisterAuthServiceServer(grpcServer, h)
	fmt.Printf("Starting server on :%v\n", port)

	go func() {
		grpcServer.Serve(lis)
	}()
	go func() {
		pubsub.Subscribe(&subscribers.AuthServiceSubscriber{Handler: h})
	}()

	<-c
	grpcServer.GracefulStop()
	pubsub.Shutdown()

}
