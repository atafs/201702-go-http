package server

import (
	"fmt"
	"net/http"
)

//Check if the server is running
func checkHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Status: I am a Running Server!!")
}

//Check if the server is running
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "I am AmericoLIB!!")
}
