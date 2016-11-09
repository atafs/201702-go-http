package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Init the server
func InitServer() {
	config, err := loadConfig()
	routes := initRoutes()
	router := newRouters(routes)

	isServerOnline := startServer(config.Port, router)
	fmt.Println("isServerOnline:", isServerOnline)

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
