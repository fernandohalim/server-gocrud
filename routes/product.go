package routes

import (
	"server-hydraulicpump/controllers/productcontroller"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	router := r.PathPrefix("/products").Subrouter()

	router.HandleFunc("", productcontroller.Index).Methods("GET")
	router.HandleFunc("", productcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}", productcontroller.Detail).Methods("GET")
	router.HandleFunc("/update/{id}", productcontroller.Update).Methods("PUT")
	router.HandleFunc("/delete/{id}", productcontroller.Delete).Methods("DELETE")
}
