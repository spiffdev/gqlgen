module github.com/spiffdev/gqlgen/_examples/large-project-structure/integration

go 1.25.0

require github.com/spiffdev/gqlgen/_examples/large-project-structure/main v0.0.0

replace github.com/spiffdev/gqlgen/_examples/large-project-structure/main => ../main

replace github.com/spiffdev/gqlgen/_examples/large-project-structure/shared => ../shared
