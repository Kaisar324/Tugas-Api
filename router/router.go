package router

import (
	"controllers"
	"middlewares"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	noteRoutes := r.PathPrefix("/notes").Subrouter()
	noteRoutes.Use(middlewares.AuthMiddleware)
	noteRoutes.HandleFunc("", controllers.CreateNote).Methods("POST")
	noteRoutes.HandleFunc("", controllers.GetNotes).Methods("GET")

	return r
}
