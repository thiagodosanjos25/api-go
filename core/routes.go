package core

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

// Router ...
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/zen", handleZen).Methods("GET")
	router.HandleFunc("/version", handleVersion).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(handleNotFound)

	return router
}
