package server

import (
	"github.com/spiffdev/gqlgen/_examples/federation/accounts/graph"
)

const DefaultPort = "4001"

var Schema = graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})
