package helpers

import (
	"encoding/json"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

func CheckCookieError(err error, w http.ResponseWriter) bool {
	if err != nil {
		if err == http.ErrNoCookie {
			json.NewEncoder(w).Encode(getMsg(http.StatusUnauthorized))
			return true
		}
		json.NewEncoder(w).Encode(getMsg(http.StatusBadRequest))
		return true
	}
	return false
}

func CheckJwtParseError(err error, w http.ResponseWriter) bool {
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			json.NewEncoder(w).Encode(getMsg(http.StatusUnauthorized))
			return true
		}
		json.NewEncoder(w).Encode(getMsg(http.StatusBadRequest))
		return true
	}
	return false
}

func CheckTokenValid(token *jwt.Token, w http.ResponseWriter) bool {
	if !token.Valid {
		json.NewEncoder(w).Encode(getMsg(http.StatusUnauthorized))
		return false
	}
	return true
}

func getMsg(httpCode int) map[string]string {
	return map[string]string{
		"message": http.StatusText(httpCode),
		"code":    strconv.FormatInt(int64(httpCode), 10),
	}
}
