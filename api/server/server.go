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
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	schemaFile, err := ioutil.ReadFile("./api/schema.graphql")
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
