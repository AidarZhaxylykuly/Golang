package main

import (
	"net/http"

<<<<<<< HEAD
	"github.com/AidarZhaxylykuly/Golang/TSIS1/pkg/handlers"
	"github.com/gorilla/mux"
=======
    "github.com/gorilla/mux"

	"github.com/AidarZhaxylykuly/Golang/cmd/handlers"
>>>>>>> 65e4009f14743a6227e0bcc1b7d18e47251d4772
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
