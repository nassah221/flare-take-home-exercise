package main

import (
	"context"
	"flag"
	"flare/exercise/data"
	"flare/exercise/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

var bindAddr = flag.String("bind", ":9090", "Bind address for the server")

func main() {
	flag.Parse()

	l := log.New(os.Stdout, "username-api ", log.LstdFlags)
	f := data.NewDB()

	ph := handlers.NewHandler(l, f)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()

	// Set up endpoints for the GET router
	getRouter.HandleFunc("/username", ph.CheckUsername).Queries("q", "{query}")
	getRouter.HandleFunc("/health", ph.CheckHealth)

	opts := middleware.RedocOpts{SpecURL: "./swagger.yml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yml", http.FileServer(http.Dir("./")))

	s := http.Server{
		Addr:         *bindAddr,         // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	// start the server
	go func() {
		l.Printf("Starting server on port %s", *bindAddr)
		l.Printf("Serving swagger docs on %s/docs", *bindAddr)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	if err := s.Shutdown(ctx); err != nil {
		l.Fatalf("[ERROR] shutting down server: %v", err)
	}
}
