package main

import (
	"log"
	"net/http"

	"github.com/eduardohumberto/test-gql/example/selection"
	"github.com/eduardohumberto/test-gql/graphql/handler"
	"github.com/eduardohumberto/test-gql/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
