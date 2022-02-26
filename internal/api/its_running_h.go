package api

import (
	"net/http"
	"strconv"

	"github.org/api-go/core"
)

// HandlerItsRunning ...
func (h *Handler) HandlerItsRunning(w http.ResponseWriter, r *http.Request) {

	response := &ResponseBodyJSON{}
	response.Code = strconv.Itoa(http.StatusOK)
	response.Message = "Success"

	core.Respond(w, r, http.StatusOK, response)
	return
}
