//go:generate go run ../../../testdata/gqlgen.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eduardohumberto/test-gql/example/federation/accounts/graph"
	"github.com/eduardohumberto/test-gql/example/federation/accounts/graph/generated"
	"github.com/eduardohumberto/test-gql/graphql/handler"
	"github.com/eduardohumberto/test-gql/graphql/handler/debug"
	"github.com/eduardohumberto/test-gql/graphql/playground"
)

const defaultPort = "4001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
