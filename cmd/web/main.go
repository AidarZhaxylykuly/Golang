package main

import (
    "net/http"

    "github.com/gorilla/mux"

	pkg "Golang/pkg/handlers"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/welcome", pkg.WelcomeMessage)
	r.HandleFunc("/health", pkg.HealthCheck).Methods("GET")
	r.HandleFunc("/DiaryMain", pkg.DiaryMain).Methods("GET")
	r.HandleFunc("/DiaryMain/{number}", pkg.DiaryDaily).Methods("GET")
	http.ListenAndServe(":8080", r)
}