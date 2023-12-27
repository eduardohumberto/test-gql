package main

import (
	"log"
	"net/http"

	"github.com/eduardohumberto/test-gql/example/scalars"
	"github.com/eduardohumberto/test-gql/graphql/handler"
	"github.com/eduardohumberto/test-gql/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", handler.NewDefaultServer(scalars.NewExecutableSchema(scalars.Config{Resolvers: &scalars.Resolver{}})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
