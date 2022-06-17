package guvnor

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Validat0rs/guvnor/pkg/guvnor/types"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type (
	Guvnor types.Guvnor
)

func NewGuvnor() *Guvnor {
	return &Guvnor{
		Secure: false,
		HTTP: types.HTTP{
			Router: mux.NewRouter(),
			Client: &http.Client{Timeout: 5 * time.Second},
		},
		Cache: types.Cache{
			Redis: redis.NewClient(&redis.Options{
				Addr:     os.Getenv("REDIS_URL"),
				Password: os.Getenv("REDIS_PASSWORD"),
				DB:       0,
			}),
		},
		Monitoring: types.Monitoring{
			Logger: log.With().Str("module", "feed").Logger(),
		},
	}
}

func (g *Guvnor) Start() {
	g.Monitoring.Logger.Info().Msgf("guvnor starting on %v....", ":"+os.Getenv("GUVNOR_PORT"))
	g.HTTP.Router.Use()
	g.HTTP.Server = &http.Server{
		Addr:         ":" + os.Getenv("GUVNOR_PORT"),
		Handler:      g.HTTP.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		if err := g.HTTP.Server.ListenAndServe(); err != nil {
			g.Monitoring.Logger.Info().Err(err)
		}
	}()
}

func (g *Guvnor) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := g.HTTP.Server.Shutdown(ctx); err != nil {
		g.Monitoring.Logger.Fatal().Err(err).Msg("guvnor shutdown")
	}

	g.Monitoring.Logger.Info().Msg("guvnor exiting....")
}
