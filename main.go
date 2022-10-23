package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	healthservice "github.com/uAccounts/uAccounts-Backend/service/healthService"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/systemhealthcheck", healthservice.HealthCheck).Methods("GET")

	handler := cors.AllowAll().Handler(r)
	fmt.Printf("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
