package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/units"
	"github.com/ben-toogood/kite/api/graph"
	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/followers"
	"github.com/ben-toogood/kite/posts"
	"github.com/ben-toogood/kite/users"
	"github.com/dgrijalva/jwt-go"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/cors"
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

	r := &graph.Resolver{
		Users:     users.NewClient(),
		Comments:  comments.NewClient(),
		Auth:      auth.NewClient(),
		Posts:     posts.NewClient(),
		Followers: followers.NewClient(),
		PublicKey: key,
	}

	gh := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: r}))

	gh.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 30 * time.Second,
	})
	gh.AddTransport(transport.Options{})
	gh.AddTransport(transport.GET{})
	gh.AddTransport(transport.POST{})
	gh.AddTransport(
		transport.MultipartForm{
			MaxUploadSize: int64(20 * units.MB),
			MaxMemory:     int64(50 * units.MB),
		},
	)
	gh.SetQueryCache(lru.New(1000))
	gh.Use(extension.Introspection{})

	h := http.Handler(gh)
	h = graph.WithLoaders(r, h)
	h = r.AuthMiddleware(h)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodPost},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	h = c.Handler(h)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", h)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
