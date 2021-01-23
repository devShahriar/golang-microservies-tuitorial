package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	goCros "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "api", log.LstdFlags)

	fileHandler := NewFile()

	sm := mux.NewRouter()
	ch := goCros.CORS(goCros.AllowedOrigins([]string{"*"}))
	server := http.Server{
		Addr:         ":9001",
		Handler:      ch(sm),
		IdleTimeout:  123 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("gracful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc) // shuts the server when users has done with the request
}
