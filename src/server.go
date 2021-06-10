package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/resurfaceio/test-mux/graph"
	"github.com/resurfaceio/test-mux/graph/generated"
)

const defaultPort = "8080"

type message struct {
	Msg string `json:"msg"`
}

func pong(w http.ResponseWriter, r *http.Request) { // handler for test ping call

	w.Header().Set("Content-Type", "application/json")
	msg := message{Msg: "pong"}
	json.NewEncoder(w).Encode(msg)
}

func LoggerMiddlewareFunc(h http.Handler) http.Handler {

	// option struct to pass to logger
	// opt := logger.Options{
	// 	Rules:   "",
	// 	Url:     "",
	// 	Enabled: true,
	// 	Queue:   make([]string, 0),
	// }

	// // create new http logger instance
	// logger, err := logger.NewHttpLogger(opt)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Whale hello there, from the middleware!", r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.HandleFunc("/ping", pong)

	router.Use(LoggerMiddlewareFunc)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
