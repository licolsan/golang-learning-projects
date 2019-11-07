package controllers

import (
	"encoding/json"
	"licolsan/constants"
	"licolsan/models"
	"licolsan/services"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds models.Credential
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users := services.GetUsers()
	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claim := &models.Claim{
		creds.Username,
		jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(os.Getenv(constants.JwtKey)))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(getMsg(tokenString))
}

func getMsg(tokenString string) map[string]string {
	return map[string]string{
		"Name":  constants.Token,
		"Value": tokenString,
	}
}
