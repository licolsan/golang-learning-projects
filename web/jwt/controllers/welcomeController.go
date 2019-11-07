package controllers

import (
	"encoding/json"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"Message": "Welcome!"})
}
