package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/luisaguilar2910/go-server/internal/app"
	"github.com/luisaguilar2910/go-server/internal/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuth)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router.HandleFunc("/api/health", controllers.HealthCheck).Methods("GET")

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
