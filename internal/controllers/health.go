package controllers

import (
	"net/http"

	u "github.com/luisaguilar2910/go-server/internal/utils"
)

var HealthCheck = func(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "The app is running well!!")
	u.Response(w, response)
	return
}
