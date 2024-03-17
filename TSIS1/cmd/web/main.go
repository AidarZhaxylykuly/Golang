package main

import (
	"net/http"

	"github.com/AidarZhaxylykuly/Golang/TSIS1/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/welcome", handlers.WelcomeMessage)
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/DiaryMain", handlers.DiaryMain).Methods("GET")
	r.HandleFunc("/DiaryMain/{number}", handlers.DiaryDaily).Methods("GET")
	http.ListenAndServe(":8080", r)
	//"github.com/AidarZhaxylykuly/Golang/cmd/handlers"
}
