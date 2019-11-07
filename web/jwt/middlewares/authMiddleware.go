package middlewares

import (
	"fmt"
	"licolsan/constants"
	"licolsan/helpers"
	"licolsan/models"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func CheckToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get(constants.Authorization)
		splitToken := strings.Split(reqToken, constants.Bearer)

		if len(splitToken) != 2 {
			fmt.Println("token is invalid")
			return
		}

		tokenString := splitToken[1]
		claim := &models.Claim{}
		token, err := jwt.ParseWithClaims(tokenString, claim, helpers.TokenValidator)

		if helpers.CheckJwtParseError(err, w) ||
			!helpers.CheckTokenValid(token, w) {
			return
		}

		next.ServeHTTP(w, r)
	})
}
