package controllers

import (
	"go-jwt/helpers"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	helpers.Response(w, 200, "My Profile", user)
}
