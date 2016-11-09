package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// The route and the handler function defined for that route.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// The array to hold all of the routes.
type Routes []Route

// Start the server
func startServer() {
	routes := initRoutes()
	router := newRouters(routes)
}

// Create the array for all of the routes that our service listens to.
func initRoutes() Routes {
	var routes = Routes{
		// Generic handlers
		Route{"Healthcheck", "GET", "/healthcheck", HealthcheckHandler},
		Route{"Ticket", "GET", "/tickets", TicketHandler},
	}

	return routes
}

//NewRouter method : Create the router for the webservice.
func newRouters(r Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
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
