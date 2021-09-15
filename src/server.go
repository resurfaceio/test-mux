package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/resurfaceio/logger-go"
	"github.com/resurfaceio/test-mux/graph"
	"github.com/resurfaceio/test-mux/graph/generated"
	database "github.com/resurfaceio/test-mux/internal/pkg/db"
)

const defaultPort = "5000"

type App struct {
	Router mux.Router
}

type message struct {
	Msg string `json:"msg"`
}

func pong(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	msg := message{
		Msg: "pong",
	}
	json.NewEncoder(w).Encode(msg)
}

func main() {
	database.InitDB()
	database.Clear()
	database.Migrate()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app := App{
		Router: *mux.NewRouter(),
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	app.Router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	app.Router.Handle("/query", srv)
	app.Router.HandleFunc("/ping", pong)

	options := logger.Options{
		Url:     "http://localhost:4001/message",
		Rules:   "allow_http_url\n/request_header:bar/ remove_if /.*foo.*/",
		Enabled: true,
		Queue:   nil,
	}

	httpLoggerForMux, err := logger.NewHttpLoggerForMuxOptions(options)
	if err != nil {
		log.Fatal(err)
	}

	app.Router.Use(httpLoggerForMux.LogData)

	log.Fatal(http.ListenAndServe(":"+port, &app.Router))
}
