package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hi there, I love %s!", req.URL.Path[1:])
}
