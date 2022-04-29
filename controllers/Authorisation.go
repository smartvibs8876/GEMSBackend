package controllers

import (
	"encoding/json"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/dgrijalva/jwt-go"
)

func Authorization(w http.ResponseWriter, r *http.Request) entity.Users {
	tokenString := r.Header.Get("Authorization")
	tokenString = tokenString[8 : len(tokenString)-1]
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(false)
		return entity.Users{}
	}
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(false)
			return entity.Users{}
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(false)
		return entity.Users{}
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(false)
		return entity.Users{}
	}
	var user entity.Users
	database.Connector.Where("email = ?", claims.Email).First(&user)
	return user
}
