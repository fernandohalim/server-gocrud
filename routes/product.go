package routes

import (
	"server-hydraulicpump/controllers/productcontroller"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	router := r.PathPrefix("/products").Subrouter()

	router.HandleFunc("", productcontroller.Index).Methods("GET")
}
