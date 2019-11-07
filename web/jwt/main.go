package main

import (
	"licolsan/controllers"
	"licolsan/initializers"
	"licolsan/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initializers.LoadEnvironment()

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/signin", controllers.Signin).Methods("POST")
	myRouter.HandleFunc("/welcome", middlewares.CheckToken(controllers.Welcome))
	myRouter.HandleFunc("/refresh", controllers.Refresh).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
