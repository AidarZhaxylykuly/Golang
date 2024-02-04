package main

import (
    "net/http"

    "github.com/gorilla/mux"

	"github.com/AidarZhaxylykuly/Golang/cmd/mypackage"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/welcome", handlers.WelcomeMessage)
	r.HandleFunc("/health", HealthCheck).Methods("GET")

	http.ListenAndServe(":8080", r)
}