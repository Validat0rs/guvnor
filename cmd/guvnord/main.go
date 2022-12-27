package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Validat0rs/guvnor/pkg/guvnor"
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

	g := guvnor.NewGuvnor()
	g.SetHandlers()
	g.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	g.Stop()
}

func checkEnv() error {
	for _, envVar := range requiredEnv {
		if os.Getenv(envVar) == "" {
			return fmt.Errorf("%s is not set", envVar)
		}
	}

	return nil
}
