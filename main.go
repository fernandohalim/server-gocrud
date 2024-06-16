package main

import (
	"fmt"
	"net/http"
	"server-hydraulicpump/config"
	"server-hydraulicpump/routes"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	log.Println("Server is running on port", config.ENV.PORT)

	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
