package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joheee/johedotcom/graph"
	"github.com/joheee/johedotcom/graph/generated"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/gorilla/websocket"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"HEAD","GET","POST","PUT","DELETE","PATCH"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		Debug: true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
        Upgrader: websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool {
                // Check against your desired domains here
                 return r.Host == "example.org"
            },
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
        },
    })
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
