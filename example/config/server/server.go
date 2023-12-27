package main

import (
	"log"
	"net/http"

	todo "github.com/eduardohumberto/test-gql/example/config"
	"github.com/eduardohumberto/test-gql/graphql/handler"
	"github.com/eduardohumberto/test-gql/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
