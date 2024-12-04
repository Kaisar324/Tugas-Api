package controllers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"

	"your_project/database"
	"your_project/models"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	var user models.User
	database.DB.Where("username = ?", username).First(&user)

	var note models.Note
	json.NewDecoder(r.Body).Decode(&note)
	note.UserID = user.ID

	database.DB.Create(&note)
	w.WriteHeader(http.StatusCreated)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	var user models.User
	database.DB.Where("username = ?", username).First(&user)

	var notes []models.Note
	database.DB.Where("user_id = ?", user.ID).Find(&notes)
	json.NewEncoder(w).Encode(notes)
}
