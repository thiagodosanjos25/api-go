package api

import (
	mux "github.com/gorilla/mux"
	newrelic "github.com/newrelic/go-agent"
	"github.org/api-go/core"
)

// Router ...
func Router(h *Handler) *mux.Router {
	router := core.Router()
	router.HandleFunc(newrelic.WrapHandleFunc(h.Relic, "/api/v1/its-running", h.HandlerItsRunning)).Methods("GET")

	// Client Entity
	router.HandleFunc(newrelic.WrapHandleFunc(h.Relic, "/api/v1/client", h.HandlerAddClient)).Methods("POST")
	router.HandleFunc(newrelic.WrapHandleFunc(h.Relic, "/api/v1/clients", h.HandlerListClients)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(h.Relic, "/api/v1/client/{id}", h.HandlerGetClient)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(h.Relic, "/api/v1/client/{id}", h.HandlerUpdateClient)).Methods("PUT")
	router.HandleFunc(newrelic.WrapHandleFunc(h.Relic, "/api/v1/client/{id}", h.HandlerDeleteClient)).Methods("DELETE")

	return router
}
