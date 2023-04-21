package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/go-chi/chi/v5"
	"github.com/prae-api/todo-api/graph"
)

const defaultPort = "8080"

func main() {
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = defaultPort
	//}
	port := defaultPort

//code-review-Jedi
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		TodoStore: map[int]model.Todo{},
	}}))
	// We should initialize the Todo map here, so the every time the resolver gets called, this Todo map is used
	// This way we don't have to check whether the map already exists inside a resolver, making the program more efficient
//code-review-Jedi

	r := chi.NewRouter()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))


	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
