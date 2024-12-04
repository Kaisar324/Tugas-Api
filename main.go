package main

import (
	"log"
	"net/http"

	"your_project/database"
	"your_project/models"
	"your_project/router"
)

func main() {
	database.InitDatabase()
	database.DB.AutoMigrate(&models.User{}, &models.Note{})

	r := router.InitRouter()

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
