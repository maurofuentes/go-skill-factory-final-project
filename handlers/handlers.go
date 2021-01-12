package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers is a function that hand routers, port and listeners
func Handlers()  {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	corsAllow := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, corsAllow))
}