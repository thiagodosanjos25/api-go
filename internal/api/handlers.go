package api

import (
	"github.org/api-go/internal/database"

	redis "gopkg.in/redis.v4"

	newrelic "github.com/newrelic/go-agent"
)

//Handler ...
type Handler struct {
	Relic        newrelic.Application
	ClientRedis  *redis.Client
	DB           *database.DataBase
	DBAux        *database.DataBase
	HostsVálidos []string
	// SnetLog      *snetlog.Log
	GodinAuthURL string
}
