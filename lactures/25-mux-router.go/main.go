package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("start mux server")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r)) // on 8080 using mux

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>this is home route.</h1>"))
}
