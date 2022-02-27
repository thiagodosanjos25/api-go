package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.org/api-go/internal/database"
)

// Client ...
type Client struct {
	Id         int     `json:"Id"`
	Name       string  `json:"name"`
	Gender     string  `json:"Gender"`
	Weight     float64 `json:"weight"`
	Height     float64 `json:"height"`
	Imc        float64 `json:"imc"`
	Situation  string  `json:"Situation"`
	Created_at string  `json:"created_at"`
	Update_at  string  `json:"update_at"`
	Active     bool    `json:"active"`
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

func (c *Client) add(h *Handler) (*Client, error) {

	if err := validateFields(c); err != nil {
		return nil, err
	}

	c.Imc, c.Situation = generateIMCandSituation(c.Weight, c.Height)

	row, err := h.DB.SelectSliceScan(sqlAddClient, nil, c.Name, c.Gender, c.Weight, c.Height, c.Imc, c.Situation)
	if err != nil {
		return nil, errors.Wrap(err, "Erro ao tentar inserir Cliente. Mensagem:")
	}

	err = sendMessageRabbitMQ(fmt.Sprintf("Cliente cadastrado com sucesso! Nome: %v", rowNil(row[0], 1)), h)
	if err != nil {
		return nil, errors.Wrap(err, "Erro ao enviar mensagem ao RabbitMQ. Mensagem:")
	}

	c.Id, _ = strconv.Atoi(rowNil(row[0], 0))
	c.Name = rowNil(row[0], 1)
	c.Gender = rowNil(row[0], 2)
	c.Weight, _ = strconv.ParseFloat(rowNil(row[0], 3), 64)
	c.Height, _ = strconv.ParseFloat(rowNil(row[0], 4), 64)
	c.Imc, _ = strconv.ParseFloat(rowNil(row[0], 5), 64)
	c.Situation = rowNil(row[0], 6)
	c.Created_at = rowNil(row[0], 7)
	c.Update_at = rowNil(row[0], 8)
	c.Active, _ = strconv.ParseBool(rowNil(row[0], 9))

	return c, err
}

func (c *Client) list(name, situation string, h *Handler) ([]*Client, error) {

	rows, err := h.DB.SelectSliceScan(sqlListClients, nil, name, name, situation, situation)
	if err != nil {
		if err == database.ErrNoRows {
			return nil, ErrNoRowsGeneric
		}
		return nil, errors.Wrap(err, "Erro ao tentar listar Clientes. Mensagem:")
	}

	clients := make([]*Client, 0, len(rows))

	for _, row := range rows {
		cAux := new(Client)

		cAux.Id, _ = strconv.Atoi(row[0].(string))
		cAux.Name = row[1].(string)
		cAux.Gender = row[2].(string)
		cAux.Weight, _ = strconv.ParseFloat(row[3].(string), 64)
		cAux.Height, _ = strconv.ParseFloat(row[4].(string), 64)
		cAux.Imc, _ = strconv.ParseFloat(row[5].(string), 64)
		cAux.Situation = row[6].(string)
		cAux.Created_at = row[7].(string)
		cAux.Update_at = row[8].(string)
		cAux.Active, _ = strconv.ParseBool(row[9].(string))

		clients = append(clients, cAux)
	}

	return clients, err
}

func (c *Client) get(idClient int, h *Handler) ([]*Client, error) {

	row, err := h.DB.SelectSliceScan(sqlGetClient, nil, idClient)
	if err != nil {
		if err == database.ErrNoRows {
			return nil, ErrNoRowsGeneric
		}
		return nil, errors.Wrap(err, "Erro ao tentar listar Cliente. Mensagem:")
	}

	cliente := make([]*Client, 0, len(row))

	c.Id, _ = strconv.Atoi(rowNil(row[0], 0))
	c.Name = rowNil(row[0], 1)
	c.Gender = rowNil(row[0], 2)
	c.Weight, _ = strconv.ParseFloat(rowNil(row[0], 3), 64)
	c.Height, _ = strconv.ParseFloat(rowNil(row[0], 4), 64)
	c.Imc, _ = strconv.ParseFloat(rowNil(row[0], 5), 64)
	c.Situation = rowNil(row[0], 6)
	c.Created_at = rowNil(row[0], 7)
	c.Update_at = rowNil(row[0], 8)
	c.Active, _ = strconv.ParseBool(rowNil(row[0], 9))

	cliente = append(cliente, c)

	return cliente, err
}

func (c *Client) update(idClient int, h *Handler) (*Client, error) {

	if err := validateFields(c); err != nil {
		return nil, err
	}

	c.Imc, c.Situation = generateIMCandSituation(c.Weight, c.Height)

	row, err := h.DB.SelectSliceScan(sqlUpdateClient, nil, c.Name, c.Gender, c.Weight, c.Height, c.Imc, c.Situation, c.Active, idClient)
	if err != nil {
		if err == database.ErrNoRows {
			return nil, ErrNoRowsGeneric
		}
		return nil, errors.Wrap(err, "Erro ao tentar editar Cliente. Mensagem:")
	}

	err = sendMessageRabbitMQ(fmt.Sprintf("Cliente editado com sucesso! Nome: %v", rowNil(row[0], 1)), h)
	if err != nil {
		return nil, errors.Wrap(err, "Erro ao enviar mensagem ao RabbitMQ. Mensagem:")
	}

	c.Id, _ = strconv.Atoi(rowNil(row[0], 0))
	c.Name = rowNil(row[0], 1)
	c.Gender = rowNil(row[0], 2)
	c.Weight, _ = strconv.ParseFloat(rowNil(row[0], 3), 64)
	c.Height, _ = strconv.ParseFloat(rowNil(row[0], 4), 64)
	c.Imc, _ = strconv.ParseFloat(rowNil(row[0], 5), 64)
	c.Situation = rowNil(row[0], 6)
	c.Created_at = rowNil(row[0], 7)
	c.Update_at = rowNil(row[0], 8)
	c.Active, _ = strconv.ParseBool(rowNil(row[0], 9))

	return c, err
}

func (c *Client) delete(idClient int, h *Handler) (*Client, error) {

	row, err := h.DB.SelectSliceScan(sqlDeleteClient, nil, idClient)
	if err != nil {
		if err == database.ErrNoRows {
			return nil, ErrNoRowsGeneric
		}
		return nil, errors.Wrap(err, "Erro ao tentar deletar Cliente. Mensagem:")
	}

	err = sendMessageRabbitMQ(fmt.Sprintf("Cliente deletado com sucesso! Nome: %v", rowNil(row[0], 1)), h)
	if err != nil {
		return nil, errors.Wrap(err, "Erro ao enviar mensagem ao RabbitMQ. Mensagem:")
	}

	c.Id, _ = strconv.Atoi(rowNil(row[0], 0))
	c.Name = rowNil(row[0], 1)
	c.Gender = rowNil(row[0], 2)
	c.Weight, _ = strconv.ParseFloat(rowNil(row[0], 3), 64)
	c.Height, _ = strconv.ParseFloat(rowNil(row[0], 4), 64)
	c.Imc, _ = strconv.ParseFloat(rowNil(row[0], 5), 64)
	c.Situation = rowNil(row[0], 6)
	c.Created_at = rowNil(row[0], 7)
	c.Update_at = rowNil(row[0], 8)
	c.Active, _ = strconv.ParseBool(rowNil(row[0], 9))

	return c, err
}

func validateFields(c *Client) error {
	if !(c.Name != "") || !(c.Gender != "") || !(c.Weight != 0) || !(c.Height != 0) {
		return ErrCamposObrigatorios
	}
	return nil
}

func generateIMCandSituation(weight, heigh float64) (imc float64, situation string) {

	imc = weight / (heigh * heigh)

	if imc < 18.5 {
		situation = "Abaixo do peso"
	} else if imc >= 18.5 && imc <= 24.9 {
		situation = "Peso normal"
	} else if imc >= 25 && imc <= 29.9 {
		situation = "Sobrepeso"
	} else if imc >= 30 && imc <= 34.9 {
		situation = "Obesidade grau 1"
	} else if imc >= 35 && imc <= 39.9 {
		situation = "Obesidade grau 2"
	} else {
		situation = "Obesidade grau 3"
	}

	return imc, situation
}

func sendMessageRabbitMQ(message string, h *Handler) error {

	channelRabbitMQ, err := h.RabbitMQ.Channel()
	if err != nil {
		return err
	}
	defer channelRabbitMQ.Close()

	_, err = channelRabbitMQ.QueueDeclare(
		"QueueClient", // queue name
		true,          // durable
		false,         // auto delete
		false,         // exclusive
		false,         // no wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}

	messageAmqp := amqp.Publishing{
		Headers:         map[string]interface{}{},
		ContentType:     "CRUD API-GO",
		ContentEncoding: "",
		DeliveryMode:    0,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       message,
		Timestamp:       time.Time{},
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            []byte(message),
	}

	if err := channelRabbitMQ.Publish(
		"",            // exchange
		"QueueClient", // queue name
		false,         // mandatory
		false,         // immediate
		messageAmqp,   // message to publish
	); err != nil {
		return err
	}

	return nil
}
