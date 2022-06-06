package proposals

import (
	"context"
	"time"
)

func (p *Proposals) getCache(cacheKey string) *string {
	feed, _ := p.redisClient.Get(context.TODO(), cacheKey).Result()

	return &feed
}

func (p *Proposals) setCache(cacheKey, feed string) error {
	_, err := p.redisClient.Set(context.TODO(), cacheKey, feed, 6*time.Hour).Result()
	if err != nil {
		return err
	}

	return nil
}
