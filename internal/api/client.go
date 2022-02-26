package api

import (
	"strconv"

	"github.com/pkg/errors"
	"github.org/api-go/internal/database"
)

// Client ...
type Client struct {
	IdMensagem        int    `json:"idMensagem"`
	IdSubRede         int    `json:"idSubRede"`
	IdEstabelecimento int    `json:"idEstabelecimento"`
	IdTerminal        int    `json:"idTerminal"`
	IdUsuario         int    `json:"idUsuario"`
	Titulo            string `json:"titulo"`
	Mensagem          string `json:"mensagem"`
	DataInicio        string `json:"dataInicio"`
	DataFim           string `json:"dataFim"`
	Ativo             bool   `json:"ativo"`
}

// RespClients ...
type RespClients struct {
	ResponseBodyJSON
	Client []*Client `json:"clients"`
}

// RespClient ...
type RespClient struct {
	ResponseBodyJSON
	Client *Client `json:"client"`
}

func (mc *Client) add(h *Handler) (*Client, error) {

	if !(mc.Titulo != "") || !(mc.Mensagem != "") || !(mc.DataInicio != "") || !(mc.DataFim != "") || !(mc.IdUsuario != 0) {
		return nil, ErrCamposObrigatorios
	}

	row, erro := h.DB.SelectSliceScan(sqlAddClient, nil, mc.IdSubRede, mc.IdSubRede, mc.IdEstabelecimento, mc.IdEstabelecimento, mc.IdTerminal, mc.IdTerminal, mc.Titulo, mc.Mensagem, mc.DataInicio, mc.DataFim, mc.Ativo, mc.IdUsuario)
	if erro != nil {
		return nil, errors.Wrap(erro, "Erro ao tentar inserir Mensagem Circular. Mensagem:")
	}

	mc.IdMensagem, _ = strconv.Atoi(rowNil(row[0], 0))

	return mc, erro
}

func (mc *Client) list(dataInicio string, dataFim string, titulo string, idSubRede int, idEstabelecimento int, idTerminal int, h *Handler) ([]*Client, error) {

	if !(dataInicio != "") || !(dataFim != "") {
		return nil, ErrCamposObrigatorios
	}

	rows, erro := h.DB.SelectSliceScan(sqlListClients, nil, dataInicio, dataFim, titulo, titulo,
		idSubRede, idSubRede, idEstabelecimento, idEstabelecimento, idTerminal, idTerminal)
	if erro != nil {
		if erro == database.ErrNoRows {
			return nil, ErrNoRowsGeneric
		}
		return nil, errors.Wrap(erro, "Erro ao tentar listar Mensagem Circular. Mensagem:")
	}

	clients := make([]*Client, 0, len(rows))

	for _, row := range rows {
		mcAux := new(Client)

		mcAux.IdMensagem, _ = strconv.Atoi(row[0].(string))
		mcAux.IdSubRede, _ = strconv.Atoi(row[1].(string))
		mcAux.IdEstabelecimento, _ = strconv.Atoi(row[2].(string))
		mcAux.IdTerminal, _ = strconv.Atoi(row[3].(string))
		mcAux.Titulo = row[4].(string)
		mcAux.Mensagem = row[5].(string)
		mcAux.DataInicio = row[6].(string)
		mcAux.DataFim = row[7].(string)
		mcAux.IdUsuario, _ = strconv.Atoi(row[8].(string))
		mcAux.Ativo, _ = strconv.ParseBool(row[9].(string))

		clients = append(clients, mcAux)
	}

	return clients, erro
}

func (mc *Client) get(idClient int, h *Handler) ([]*Client, error) {

	row, erro := h.DB.SelectSliceScan(sqlGetClient, nil, idClient)
	if erro != nil {
		if erro == database.ErrNoRows {
			return nil, ErrNoRowsGeneric
		}
		return nil, errors.Wrap(erro, "Erro ao tentar listar Mensagem Circular. Mensagem:")
	}

	mensagemCircular := make([]*Client, 0, len(row))

	mc.IdMensagem, _ = strconv.Atoi(rowNil(row[0], 0))
	mc.IdSubRede, _ = strconv.Atoi(rowNil(row[0], 1))
	mc.IdEstabelecimento, _ = strconv.Atoi(rowNil(row[0], 2))
	mc.IdTerminal, _ = strconv.Atoi(rowNil(row[0], 3))
	mc.Titulo = rowNil(row[0], 4)
	mc.Mensagem = rowNil(row[0], 5)
	mc.DataInicio = rowNil(row[0], 6)
	mc.DataFim = rowNil(row[0], 7)
	mc.IdUsuario, _ = strconv.Atoi(rowNil(row[0], 8))
	mc.Ativo, _ = strconv.ParseBool(rowNil(row[0], 9))

	mensagemCircular = append(mensagemCircular, mc)

	return mensagemCircular, erro
}

func (mc *Client) update(idMensagem int, h *Handler) (*Client, error) {

	if !(mc.Titulo != "") || !(mc.Mensagem != "") || !(mc.DataInicio != "") || !(mc.DataFim != "") || !(mc.IdUsuario != 0) {
		return nil, ErrCamposObrigatorios
	}

	row, erro := h.DB.SelectSliceScan(sqlUpdateClient, nil, mc.IdSubRede, mc.IdSubRede, mc.IdEstabelecimento, mc.IdEstabelecimento, mc.IdTerminal, mc.IdTerminal,
		mc.Titulo, mc.Mensagem, mc.DataInicio, mc.DataFim, mc.IdUsuario, mc.Ativo, idMensagem)
	if erro != nil {
		if erro == database.ErrNoRows {
			return nil, ErrNoRowsGeneric
		}
		return nil, errors.Wrap(erro, "Erro ao tentar editar Mensagem Circular. Mensagem:")
	}

	mc.IdMensagem, _ = strconv.Atoi(rowNil(row[0], 0))

	return mc, erro
}

func (mc *Client) delete(idMensagem int, h *Handler) (*Client, error) {

	row, erro := h.DB.SelectSliceScan(sqlDeleteClient, nil, idMensagem)
	if erro != nil {
		if erro == database.ErrNoRows {
			return nil, ErrNoRowsGeneric
		}
		return nil, errors.Wrap(erro, "Erro ao tentar deletar Mensagem Circular. Mensagem:")
	}

	mc.IdMensagem, _ = strconv.Atoi(rowNil(row[0], 0))
	mc.IdSubRede, _ = strconv.Atoi(rowNil(row[0], 1))
	mc.IdEstabelecimento, _ = strconv.Atoi(rowNil(row[0], 2))
	mc.IdTerminal, _ = strconv.Atoi(rowNil(row[0], 3))
	mc.Titulo = rowNil(row[0], 4)
	mc.Mensagem = rowNil(row[0], 5)
	mc.DataInicio = rowNil(row[0], 6)
	mc.DataFim = rowNil(row[0], 7)
	mc.IdUsuario, _ = strconv.Atoi(rowNil(row[0], 8))
	mc.Ativo, _ = strconv.ParseBool(rowNil(row[0], 9))

	return mc, erro
}
