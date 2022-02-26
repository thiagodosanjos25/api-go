package api

import "errors"

var (
	ErrDecodeJson         = errors.New("Erro no formato do Json do formulario")
	ErrCamposObrigatorios = errors.New("Parâmetros obrigatórios não preenchidos")
	ErrNoRowsGeneric      = errors.New("Não houve registros no retorno da query")
)
