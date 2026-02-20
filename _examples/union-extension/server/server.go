package main

import (
	"log"
	"net/http"

	unionextension "github.com/spiffdev/gqlgen/_examples/union-extension"
	"github.com/spiffdev/gqlgen/graphql/handler"
	"github.com/spiffdev/gqlgen/graphql/handler/transport"
	"github.com/spiffdev/gqlgen/graphql/playground"
)

func main() {
	srv := handler.New(
		unionextension.NewExecutableSchema(
			unionextension.Config{Resolvers: &unionextension.Resolver{}},
		),
	)
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	http.Handle("/", playground.Handler("Union Extension Demo", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":8086", nil))
}
