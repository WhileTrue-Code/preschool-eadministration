package controller

import (
	"authorization"
	"encoding/json"
	"github.com/cristalhq/jwt/v4"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getIDFromReqAsPrimitive(writer http.ResponseWriter, req *http.Request) (primitive.ObjectID, error) {
	vars := mux.Vars(req)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return primitive.NilObjectID, err
	}

	return objectID, nil
}
