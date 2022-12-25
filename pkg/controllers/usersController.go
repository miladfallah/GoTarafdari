package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/miladfallah/GoTarafdari/pkg/models"
	"github.com/miladfallah/GoTarafdari/pkg/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUsers := &models.Users{}
	utils.ParseBody(r, CreateUsers)
	b := CreateUsers.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
