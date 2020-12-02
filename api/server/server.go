package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ben-toogood/kite/api/resolvers"
	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/users"
	"github.com/friendsofgo/graphiql"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/opentracing/opentracing-go"
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

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.MaxParallelism(20)}
	schema := graphql.MustParseSchema(string(schemaFile), &resolvers.Resolver{
		Users:    users.NewClient(),
		Comments: comments.NewClient(),
	}, opts...)

	http.Handle("/", graphiqlHandler)
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	logrus.Info("GraphQL API started on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
