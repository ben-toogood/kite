package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ben-toogood/kite/api/resolvers"
	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/followers"
	"github.com/ben-toogood/kite/posts"
	"github.com/ben-toogood/kite/users"
	"github.com/form3tech-oss/jwt-go"
	"github.com/friendsofgo/graphiql"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Load the certs
	keyPath := os.Getenv("PUBLIC_KEY_FILEPATH")
	if len(keyPath) == 0 {
		panic("Missing PUBLIC_KEY_FILEPATH")
	}
	file, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(file)
	if err != nil {
		panic(err)
	}

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

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	schemaFile, err := ioutil.ReadFile("./schema.graphql")
	if err != nil {
		log.Fatal(err)
	}

	r := &resolvers.Resolver{
		Users:     users.NewClient(),
		Comments:  comments.NewClient(),
		Auth:      auth.NewClient(),
		Posts:     posts.NewClient(),
		Followers: followers.NewClient(),
		PublicKey: key,
	}
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.MaxParallelism(20)}
	schema := graphql.MustParseSchema(string(schemaFile), r, opts...)

	mux := http.NewServeMux()

	mux.Handle("/", graphiqlHandler)
	mux.Handle("/graphql", nethttp.Middleware(tracer, r.AuthMiddleware(resolvers.WithLoaders(r, &relay.Handler{Schema: schema}))))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodPost},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	logrus.Info("GraphQL API started on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(mux)))
}
