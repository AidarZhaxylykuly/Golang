package handlers

import (
		"net/http"
		"fmt"
		"encoding/json"
		"github.com/gorilla/mux"
		"strconv"
)


func WelcomeMessage(w http.ResponseWriter, r *http.Request){
	fmt.Println(w, "Welcome! That's my weekly diary of tasks, those I should do")
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "It's ok")
}

func DiaryMain(w http.ResponseWriter, r *http.Request){
	days:=AllWeeklyTasks()
	json.NewDecoder(w).Encode(days)
}

func DiaryDaily(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["number"])
	if err != nil {
		return
	}
	player, found := BytheDay(num)
	if found {
		json.NewEncoder(w).Encode(player)
	} else {
		fmt.Fprintf(w, "Such a lot of days in one week?")
	}
	
}