package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/jaslife1/shippingcost-server/graph"
	"github.com/jaslife1/shippingcost-server/graph/generated"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/gorilla/websocket"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

    // Add CORS middleware around every request
    // See https://github.com/rs/cors for full option listing
    // router.Use(cors.New(cors.Options{
    //     AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
    //     AllowOriginFunc:  func(origin string) bool { return true },
    //     AllowedMethods:   []string{},
    //     AllowedHeaders:   []string{},
    //     AllowCredentials: true,
    //     Debug:            true,
    // }).Handler)
	router.Use(cors.Default().Handler)

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
    srv.AddTransport(&transport.Websocket{
        Upgrader: websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool {
                // Check against your desired domains here
                return r.Host == "localhost:8080"
            },
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
        },
    })

    //router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/", srv)
    router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    err := http.ListenAndServe(":"+port, router)
    if err != nil {
        panic(err)
    }
}
