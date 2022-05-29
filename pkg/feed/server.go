package feed

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Validat0rs/guvnor/pkg/feed/types"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type (
	Feed types.Feed
)

func NewFeed() *Feed {
	return &Feed{
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

func (f *Feed) Start() {
	f.Monitoring.Logger.Info().Msgf("guvnor feed service starting on %v....", ":"+os.Getenv("GUVNOR_PORT"))
	f.HTTP.Router.Use()
	f.HTTP.Server = &http.Server{
		Addr:         ":" + os.Getenv("GUVNOR_PORT"),
		Handler:      f.HTTP.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		if err := f.HTTP.Server.ListenAndServe(); err != nil {
			f.Monitoring.Logger.Info().Err(err)
		}
	}()
}

func (f *Feed) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := f.HTTP.Server.Shutdown(ctx); err != nil {
		f.Monitoring.Logger.Fatal().Err(err).Msg("guvnor shutdown")
	}

	f.Monitoring.Logger.Info().Msg("guvnor exiting....")
}
