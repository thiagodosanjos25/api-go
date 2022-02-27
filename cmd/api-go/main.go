package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	redis "github.com/go-redis/redis"
	newrelic "github.com/newrelic/go-agent"
	"github.com/rs/cors"
	"github.com/streadway/amqp"
	"github.org/api-go/config"
	"github.org/api-go/internal/api"
	"github.org/api-go/internal/database"
)

const (
	//TimeOutSecond ...
	TimeOutSecond = 120
)

var (
	arqConfig string
)

func init() {
	flag.StringVar(&arqConfig, "conf", "", "Arquivo de configuracoes em formato json")
}

func main() {
	flag.Parse()
	log.Println("Api-Go (Versao: prd_lts)")
	config := config.NewConfig(arqConfig)

	configNewRelic := newrelic.NewConfig("api-go", config.NewRelicToken)
	app, err := newrelic.NewApplication(configNewRelic)
	if err != nil {
		log.Println("Erro ao iniciar o New Relic. Erro:", err)
	}

	log.Printf("Tentando conectar no Redis...")
	redisOption, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		log.Println("Erro no ParseURL cache redis. Erro=", err)
		panic("Cache Redis não foi iniciado com sucesso.")
	}
	clientRedis := redis.NewClient(redisOption)
	_, err = clientRedis.Ping().Result()
	if err != nil {
		log.Println("Erro ao iniciar o cache redis. Erro=", err)
		panic("Cache Redis não foi iniciado com sucesso.")
	}

	log.Printf("Conectado Redis.\n")

	// connectRabbitMQ, err := amqp.Dial("amqp://guest:guest@message-queue:5672/")
	connectRabbitMQ, err := amqp.Dial(config.RabbitMQUrl)
	if err != nil {
		log.Printf("Erro ao conectar no RabbitMQ. Mensagem:")
	}

	defer connectRabbitMQ.Close()

	optionDB := &database.OptionsDB{DriverName: "postgres", IP: config.DBBizHost, Porta: config.DBBizPorta,
		NomeDB: config.DBBizNome, User: config.DBBizUser, Senha: config.DBBizSenha, Debug: false, Alias: config.DBBizNome,
		TamPoolIdleConn: 1, TempoPoolIdleConn: 1, LogMinDuration: 1000}

	db := database.NewDB(optionDB)
	if err = db.Open(); err != nil {
		log.Println("Erro ao conectar no DB. Erro=", err)
	} else {
		fmt.Printf("Database conectado!.\n")
	}
	defer db.Close()

	allowedParam := make(map[string][]string)
	if err := json.Unmarshal([]byte(config.AllowedParam), &allowedParam); err != nil {
		log.Println("Erro no json Unmarshal do allowedOrigins. Detalhe:", err)
		os.Exit(1)
	}

	c := cors.New(cors.Options{
		AllowedOrigins: allowedParam["Origins"],
		AllowedHeaders: allowedParam["Headers"],
		AllowedMethods: allowedParam["Methods"],
		Debug:          false,
	})

	hapi := &api.Handler{Relic: app, DB: db, RabbitMQ: connectRabbitMQ}

	router := api.Router(hapi)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      c.Handler(router),
		ReadTimeout:  TimeOutSecond * time.Second,
		WriteTimeout: TimeOutSecond * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println("Erro ao iniciar Server. Erro:", err)
		os.Exit(1)
	}
}
