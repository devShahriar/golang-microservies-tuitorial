package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("hellp")
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Opsss", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, req *http.Request) {
		log.Println("hello")
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Opsss", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello%s", data)
	})

	http.ListenAndServe(":9000", nil)

}
