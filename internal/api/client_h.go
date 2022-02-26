package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
	"github.org/api-go/core"
)

// HandlerAddClient ...
func (h *Handler) HandlerAddClient(w http.ResponseWriter, r *http.Request) {
	mc := &Client{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&mc); err != nil {
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(), Code: strconv.Itoa(http.StatusBadRequest),
				Message: "Erro no formato do Json"})
		return
	}

	obj, err := mc.add(h)
	if err != nil {
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(), Code: strconv.Itoa(http.StatusBadRequest),
				Message: "Erro na função"})
		return
	}

	response := &RespClient{Client: obj}
	response.Code = strconv.Itoa(http.StatusCreated)
	response.Message = "OK"

	core.Respond(w, r, http.StatusCreated, response)
	return
}

// HandlerListClient ...
func (h *Handler) HandlerListClients(w http.ResponseWriter, r *http.Request) {
	mc := &Client{}
	dataInicio := r.URL.Query().Get("dataInicio")
	dataFim := r.URL.Query().Get("dataFim")
	titulo := r.URL.Query().Get("titulo")
	idSubRede, _ := strconv.Atoi(r.URL.Query().Get("idSubRede"))
	idEstabelecimento, _ := strconv.Atoi(r.URL.Query().Get("idEstabelecimento"))
	idTerminal, _ := strconv.Atoi(r.URL.Query().Get("idTerminal"))

	obj, err := mc.list(dataInicio, dataFim, titulo, idSubRede, idEstabelecimento, idTerminal, h)
	if err != nil {
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(),
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Erro na função"})
		return
	}

	response := &RespClients{Client: obj}
	response.Code = strconv.Itoa(http.StatusOK)
	response.Message = "OK"

	core.Respond(w, r, http.StatusOK, response)
	return
}

// HandlerGetClient ...
func (h *Handler) HandlerGetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idMensagem, _ := strconv.Atoi(vars["id"])

	mc := &Client{}

	obj, err := mc.get(idMensagem, h)
	if err != nil {
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(),
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Erro na função"})
		return
	}

	response := &RespClients{Client: obj}
	response.Code = strconv.Itoa(http.StatusOK)
	response.Message = "OK"

	core.Respond(w, r, http.StatusOK, response)
	return
}

// HandlerUpdateClient ...
func (h *Handler) HandlerUpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idMensagem, _ := strconv.Atoi(vars["id"])

	mc := &Client{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&mc); err != nil {
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(), Code: strconv.Itoa(http.StatusBadRequest),
				Message: "Erro no formato do Json"})
		return
	}

	obj, err := mc.update(idMensagem, h)
	if err != nil {
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(), Code: strconv.Itoa(http.StatusBadRequest),
				Message: "Erro na função"})
		return
	}

	response := &RespClient{Client: obj}
	response.Code = strconv.Itoa(http.StatusOK)
	response.Message = "OK"

	core.Respond(w, r, http.StatusOK, response)
	return
}

// HandleDeleteClient ...
func (h *Handler) HandlerDeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idMensagem, _ := strconv.Atoi(vars["id"])

	mc := &Client{}

	obj, err := mc.delete(idMensagem, h)
	if err != nil {
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(), Code: strconv.Itoa(http.StatusBadRequest),
				Message: "Erro na função"})
		return
	}

	response := &RespClient{Client: obj}
	response.Code = strconv.Itoa(http.StatusOK)
	response.Message = "OK"

	core.Respond(w, r, http.StatusOK, response)
	return
}
