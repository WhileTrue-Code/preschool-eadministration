package controller

import (
	"authorization"
	"encoding/json"
	"github.com/cristalhq/jwt/v4"
	"log"
	"net/http"
	"strings"
)

func jsonResponse(object interface{}, w http.ResponseWriter) {
	resp, err := json.Marshal(object)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func extractJMBGFromClaims(writer http.ResponseWriter, req *http.Request) (string, error) {
	bearer := req.Header.Get("Authorization")
	bearerToken := strings.Split(bearer, "Bearer ")
	tokenString := bearerToken[1]

	token, err := jwt.Parse([]byte(tokenString), verifier)

	if err != nil {
		log.Println(err)
		http.Error(writer, "unauthorized", http.StatusUnauthorized)
		return "", err
	}

	claims := authorization.GetMapClaims(token.Bytes())
	jmbg := claims["jmbg"]
	return jmbg, nil
}
