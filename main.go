package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mico/handlers"
)

func main() {

	l := log.New(os.Stdout, "api", log.LstdFlags)
	handler := handlers.NewRouter(l)

	sh := http.NewServeMux()
	sh.Handle("/", handler)
	http.ListenAndServe(":9000", sh)

}
