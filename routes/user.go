package routes

import (
	"server-hydraulicpump/controllers/usercontroller"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", usercontroller.Index).Methods("GET")
}
