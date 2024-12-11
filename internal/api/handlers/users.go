package handlers

import (
	"encoding/json"
	"flyte/goapi/internal/api/services"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
