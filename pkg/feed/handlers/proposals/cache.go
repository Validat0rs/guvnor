package proposals

import (
	"context"
	"time"
)

func (p *Proposals) getCache() *string {
	feed, _ := p.redisClient.Get(context.TODO(), p.chainId).Result()

	return &feed
}

func (p *Proposals) setCache(feed string) error {
	_, err := p.redisClient.Set(context.TODO(), p.chainId, feed, 6*time.Hour).Result()
	if err != nil {
		return err
	}

	return nil
}
