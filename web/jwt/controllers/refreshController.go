package controllers

import (
	"licolsan/constants"
	"licolsan/helpers"
	"licolsan/models"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(constants.Token)

	if helpers.CheckCookieError(err, w) {
		return
	}

	tokenString := cookie.Value
	claim := &models.Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claim, helpers.TokenValidator)

	if helpers.CheckJwtParseError(err, w) ||
		!helpers.CheckTokenValid(token, w) {
		return
	}

	if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claim.ExpiresAt = expirationTime.Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString([]byte(os.Getenv(constants.JwtKey)))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    constants.Token,
		Value:   tokenString,
		Expires: expirationTime,
	})
}
