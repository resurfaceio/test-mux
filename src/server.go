package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/resurfaceio/logger-go/v3"
	"github.com/resurfaceio/test-mux/graph"
	"github.com/resurfaceio/test-mux/graph/generated"
	database "github.com/resurfaceio/test-mux/internal/pkg/db"
)

const defaultPort = "5000"

var count int

type App struct {
	Router mux.Router
}

type message struct {
	Msg string `json:"msg"`
}

func pongInc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	count += 1
	msg := []byte(fmt.Sprint(count))
	w.Write(msg)
}

func pong(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	msg := message{
		Msg: "pong",
	}
	json.NewEncoder(w).Encode(msg)
}

func main() {
	count = 0
	database.InitDB()
	// database.Clear()
	database.Migrate()

	database.Populate()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app := App{
		Router: *mux.NewRouter(),
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	app.Router.Handle("/", srv)
	app.Router.Handle("/query", playground.Handler("GraphQL playground", "/"))
	app.Router.HandleFunc("/ping", pong)
	app.Router.HandleFunc("/pingInc", pongInc)

	options := logger.Options{
		Url:     "http://localhost:7701/message",
		Rules:   "allow_http_url",
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
