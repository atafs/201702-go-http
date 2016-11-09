package server

import (
	"fmt"
	"net/http"
)

//HealthcheckHandler Method to check that the service is running.
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	DisplayDebugMessage(r.Method, r.RequestURI)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Status: Fabulous")
}

func TicketHandler(w http.ResponseWriter, r *http.Request) {
	DisplayDebugMessage(r.Method, r.RequestURI)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Status: Fabulous and More!!")
}
