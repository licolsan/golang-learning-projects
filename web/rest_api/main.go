package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	articleHandler "./handlers/article"
	homeHandler "./handlers/home"
)

func main() {
	handleRequest()
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// Welcome page
	myRouter.HandleFunc("/", homeHandler.Welcome)

	// Article resource
	myRouter.HandleFunc("/articles", articleHandler.All)
	myRouter.HandleFunc("/article", articleHandler.Create).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", articleHandler.Get)
	myRouter.HandleFunc("/article/{id}", articleHandler.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
