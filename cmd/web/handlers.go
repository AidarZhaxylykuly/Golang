package handlers

import ("net/http"
		"fmt"
)


func WelcomeMessage(w http.ResponseWriter, r *http.Request){
	fmt.Println(w, "Welcome")
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "It's ok")
}
