package helpers

import (
	"licolsan/constants"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func TokenValidator(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv(constants.JwtKey)), nil
}
