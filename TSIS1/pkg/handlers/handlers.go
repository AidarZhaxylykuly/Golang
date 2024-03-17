package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AidarZhaxylykuly/Golang/TSIS1/pkg/models"
	"github.com/gorilla/mux"
)

func WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome! That's my weekly diary of tasks, those I should do")


func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "It's ok")
}

func DiaryMain(w http.ResponseWriter, r *http.Request) {
	days := models.AllWeeklyTasks()
	err := json.NewEncoder(w).Encode(days)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DiaryDaily(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["number"])
	if err != nil {
		return
	}
	player, found := models.BytheDay(num)
	if found {
		json.NewEncoder(w).Encode(player)
	} else {
		fmt.Fprintf(w, "Such a lot of days in one week?")
	}

}
