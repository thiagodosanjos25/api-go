package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/caarlos0/env"
)

var config *Configuracoes

// Configuracoes ...
type Configuracoes struct {
	DBBizNome     string `json:"DBBizNome" env:"DB_NAME"`
	DBBizHost     string `json:"DBBizHost" env:"DB_HOST"`
	DBBizPorta    int    `json:"DBBizPorta" env:"DB_PORT"`
	DBBizUser     string `json:"DBBizUser" env:"DB_USER"`
	DBBizSenha    string `json:"DBBizSenha" env:"DB_PSWD"`
	EnableLogFile bool   `json:"enableLogFile" env:"ENABLE_LOG_FILE"`
	LogFile       string `json:"logFile" env:"LOG_FILE"`
	RedisHost     string `json:"redisHost" env:"REDIS_HOST"`
	RedisSenha    string `json:"redisSenha" env:"REDIS_PSWD"`
	NewRelicToken string `json:"newRelicToken" env:"NEWRELIC_TOKEN"`
	Port          int    `json:"port" env:"PORT"`
	AllowedParam  string `json:"allowedParam" env:"ALLOWED_PARAM"`
}

// NewConfig ...
func NewConfig(file string) *Configuracoes {
	var erro error

	conf := &Configuracoes{}

	if file != "" {
		fmt.Println(file)

		bufConf, err := ioutil.ReadFile(file)
		if err == nil {
			erro = json.Unmarshal(bufConf, conf)
			if erro != nil {
				log.Println(erro)
			}
		}
	}

	// variaveis de ambiente sobrescrevem informacoes do json
	if erro = env.Parse(conf); erro != nil {
		log.Println(erro)
	}

	config = conf
	return conf
}

// Config ...
func Config() *Configuracoes {
	return config
}
