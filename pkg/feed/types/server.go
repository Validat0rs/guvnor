package types

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type HTTP struct {
	Router *mux.Router
	Server *http.Server
	Client *http.Client
}

type Cache struct {
	Redis *redis.Client
}

type Monitoring struct {
	Logger zerolog.Logger
}

type Feed struct {
	Secure     bool
	HTTP       HTTP
	Cache      Cache
	Monitoring Monitoring
}
