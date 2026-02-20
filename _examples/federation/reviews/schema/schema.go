package schema

import (
	"github.com/spiffdev/gqlgen/_examples/federation/reviews/graph"
)

const DefaultPort = "4003"

var Schema = graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})
