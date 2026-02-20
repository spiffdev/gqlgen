package main

import (
	"log"
	"net/http"

	todo "github.com/spiffdev/gqlgen/_examples/config"
	"github.com/spiffdev/gqlgen/graphql/handler"
	"github.com/spiffdev/gqlgen/graphql/handler/transport"
	"github.com/spiffdev/gqlgen/graphql/playground"
)

func main() {
	srv := handler.New(
		todo.NewExecutableSchema(todo.New()),
	)
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
