package main

import (
	"log"
	"net/http"

	"github.com/spiffdev/gqlgen/_examples/selection"
	"github.com/spiffdev/gqlgen/graphql/handler"
	"github.com/spiffdev/gqlgen/graphql/handler/transport"
	"github.com/spiffdev/gqlgen/graphql/playground"
)

func main() {
	srv := handler.New(
		selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}}),
	)
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":8086", nil))
}
