package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/resurfaceio/logger-go/src/logger"
	"github.com/resurfaceio/test-mux/graph"
	"github.com/resurfaceio/test-mux/graph/generated"
)

const defaultPort = "8080"

type App struct {
	Router mux.Router
	// DB, WIP add database to app
}

type message struct {
	Msg string `json:"msg"`
}

func pong(w http.ResponseWriter, r *http.Request) { // handler for test ping call

	w.Header().Set("Content-Type", "application/json")
	msg := message{
		Msg: "pong",
	}
	json.NewEncoder(w).Encode(msg)
}

func LoggerMiddlewareFunc(next http.Handler) http.Handler {

	// option struct to pass to logger
	// opt := logger.Options{
	// 	Rules:   "allow_http_url",
	// 	Url:     "http://resurface:4001/message",
	// 	Enabled: true,
	// 	Queue:   nil,
	// }

	// create new logger instance
	// httpLoggerForMux, err := logger.NewHttpLoggerForMuxOptions(opt)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// before

		next.ServeHTTP(w, r)
		// after
		log.Println(r.Host, r.Method, r.URL, r.RequestURI)
	})
}

func main() {
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

	app.Router.Use(logger.Log) //WIP

	log.Fatal(http.ListenAndServe(":"+port, &app.Router))
}
