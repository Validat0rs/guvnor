package proposals

import (
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
)

type Proposals struct {
	chainId     string
	api         string
	redisClient *redis.Client
	logger      zerolog.Logger
}

func NewProposals(chainId, api string, redisClient *redis.Client, log zerolog.Logger) *Proposals {
	return &Proposals{
		chainId:     chainId,
		api:         api,
		redisClient: redisClient,
		logger:      log,
	}
}
