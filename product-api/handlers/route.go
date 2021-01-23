package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Router exported
type Router struct {
	l *log.Logger
}

//NewRouter exported
func NewRouter(l *log.Logger) *Router {
	return &Router{l}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("hellp")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Opsss", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", data)

}
