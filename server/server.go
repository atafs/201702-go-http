package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Init the server
func initServer() (conf Config, router *mux.Router) {
	config, err := config.loadConfig()
	routes := initRoutes()
	router := newRouters(routes)

	return config, router
}

// Start the server
func startServer(port string, router *mux.Router) bool {

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Error in the star of the server: isServerOnline=", false)
		return false
	}
	fmt.Println("Server starting post " + port + ". Listening.....")
	return true
}
