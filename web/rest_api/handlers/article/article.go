package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	dbService "../../services/db"
)

// All -
func All(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: article.All")
	json.NewEncoder(w).Encode(dbService.AllArticle())
}

// Get -
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("Endpoint Hit: article.Get")
	json.NewEncoder(w).Encode(dbService.GetArticle(id))
}

// Create -
func Create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	article := dbService.CreateArticle(reqBody)

	fmt.Println("Endpoint Hit: article.Create")
	json.NewEncoder(w).Encode(article)
}

// Delete -
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("Endpoint Hit: article.Delete")
	json.NewEncoder(w).Encode(dbService.DeleteArticle(id))
}
