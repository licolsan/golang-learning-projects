package home

import (
	"fmt"
	"net/http"
)

// Welcome -
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homePage")
}
