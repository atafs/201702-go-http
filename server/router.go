package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/web-app-go/logger"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Global variables
var (
	router *mux.Router
)

func InitServer() {
	// Initalize routes.
	routes := InitRoutes()
	toStringRoutes(routes)

	// Create server with routes
	routerReceived := NewRouter(routes)
	router = routerReceived
}

//InitRoutes : Create the array for all of the routes that our service listens to.
func InitRoutes() Routes {
	var routes = Routes{
		Route{"Index", "GET", "/", indexHandler},
		Route{"Check", "GET", "/check", checkHandler},
	}

	return routes
}

func NewRouter(r Routes) *mux.Router {
	router = mux.NewRouter().StrictSlash(true)

	for _, route := range r {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func StartServer(portToUse string) {
	logger.GetNotice("Server starting... port " + portToUse)

	err := http.ListenAndServe(":"+portToUse, router)
	if err != nil {
		log.Println("Error Starting server:", err)
	}
}

// Print
func toStringRoutes(routes Routes) {
	logger.GetNotice("Success in creating the Routes:")
	for _, route := range routes {
		logger.GetDebugf("\t=> Name:", route.Name)
		logger.GetDebugf("\t=> Method:", route.Method)
		logger.GetDebugf("\t=> Pattern:", route.Pattern)
		log.Println("\t=> HandlerFunc:", route.HandlerFunc)
		fmt.Println("")
	}
}
