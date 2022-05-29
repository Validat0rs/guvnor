package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Validat0rs/guvnor/pkg/feed"
)

var requiredEnv = []string{
	"GUVNOR_PORT",
	"GUVNOR_CONFIG",
	"GUVNOR_DOMAIN",
	"REDIS_URL",
	"PROPOSAL_FEED_AUTHOR",
	"PROPOSAL_FEED_EMAIL",
}

func main() {
	if err := checkEnv(); err != nil {
		log.Fatal(err)
	}

	_feed := feed.NewFeed()
	_feed.SetHandlers()
	_feed.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	_feed.Stop()
}

func checkEnv() error {
	for _, envVar := range requiredEnv {
		if os.Getenv(envVar) == "" {
			return fmt.Errorf("%s is not set", envVar)
		}
	}

	return nil
}
