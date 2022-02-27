package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
	"github.org/api-go/core"
)

// HandlerAddClient ...
func (h *Handler) HandlerAddClient(w http.ResponseWriter, r *http.Request) {
	c := &Client{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		log.Println(fmt.Sprintf("Erro no formato do Json.  Mensagem: %v", err.Error()))
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(), Code: strconv.Itoa(http.StatusBadRequest),
				Message: "Erro no formato do Json"})
		return
	}

	obj, err := c.add(h)
	if err != nil {
		log.Println(fmt.Sprintf("Erro na função. Mensagem: %v", err.Error()))
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(), Code: strconv.Itoa(http.StatusBadRequest),
				Message: "Erro na função"})
		return
	}

	response := &RespClient{Client: obj}
	response.Code = strconv.Itoa(http.StatusCreated)
	response.Message = "Sucesso"

	core.Respond(w, r, http.StatusCreated, response)
	return
}

// HandlerListClient ...
func (h *Handler) HandlerListClients(w http.ResponseWriter, r *http.Request) {
	c := &Client{}
	name := r.URL.Query().Get("name")
	situation := r.URL.Query().Get("situation")

	obj, err := c.list(name, situation, h)
	if err != nil {
		log.Println(fmt.Sprintf("Erro na função. Mensagem: %v", err.Error()))
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(),
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Erro na função"})
		return
	}

	response := &RespClients{Client: obj}
	response.Code = strconv.Itoa(http.StatusOK)
	response.Message = "Sucesso"

	core.Respond(w, r, http.StatusOK, response)
	return
}

// HandlerGetClient ...
func (h *Handler) HandlerGetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient, _ := strconv.Atoi(vars["id"])

	c := &Client{}

	obj, err := c.get(idClient, h)
	if err != nil {
		log.Println(fmt.Sprintf("Erro na função. Mensagem: %v", err.Error()))
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(),
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Erro na função"})
		return
	}

	response := &RespClients{Client: obj}
	response.Code = strconv.Itoa(http.StatusOK)
	response.Message = "Sucesso"

	core.Respond(w, r, http.StatusOK, response)
	return
}

// HandlerUpdateClient ...
func (h *Handler) HandlerUpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient, _ := strconv.Atoi(vars["id"])

	c := &Client{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		log.Println(fmt.Sprintf("Erro no formato do Json.  Mensagem: %v", err.Error()))
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(), Code: strconv.Itoa(http.StatusBadRequest),
				Message: "Erro no formato do Json"})
		return
	}

	obj, err := c.update(idClient, h)
	if err != nil {
		log.Println(fmt.Sprintf("Erro na função. Mensagem: %v", err.Error()))
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(), Code: strconv.Itoa(http.StatusBadRequest),
				Message: "Erro na função"})
		return
	}

	response := &RespClient{Client: obj}
	response.Code = strconv.Itoa(http.StatusOK)
	response.Message = "Sucesso"

	core.Respond(w, r, http.StatusOK, response)
	return
}

// HandleDeleteClient ...
func (h *Handler) HandlerDeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient, _ := strconv.Atoi(vars["id"])

	c := &Client{}

	obj, err := c.delete(idClient, h)
	if err != nil {
		log.Println(fmt.Sprintf("Erro na função. Mensagem: %v", err.Error()))
		core.RespondErro(w, r, http.StatusBadRequest,
			&core.ErrMessage{Erro: err.Error(), Code: strconv.Itoa(http.StatusBadRequest),
				Message: "Erro na função"})
		return
	}

	response := &RespClient{Client: obj}
	response.Code = strconv.Itoa(http.StatusOK)
	response.Message = "Sucesso"

	core.Respond(w, r, http.StatusOK, response)
	return
}
