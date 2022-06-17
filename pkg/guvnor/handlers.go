package guvnor

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Validat0rs/guvnor/pkg/guvnor/handlers/proposals"
	"github.com/Validat0rs/guvnor/pkg/guvnor/types"

	"github.com/urfave/negroni"
)

func (g *Guvnor) SetHandlers() {
	g.proposalsHandler()
	g.statusHandler()
}

func (g *Guvnor) proposalsHandler() {
	config, err := g.ParseConfig(&Reader{fileName: os.Getenv("GUVNOR_CONFIG")})
	if err != nil {
		log.Fatal(err)
	}

	for _, chain := range config.Chain {
		_proposals := proposals.NewProposals(chain.ChainId, chain.API, g.Cache.Redis, g.Monitoring.Logger)
		g.HTTP.Router.Handle("/"+chain.ChainId+"/proposals/rss", negroni.New(
			negroni.Wrap(http.HandlerFunc(_proposals.Rss)),
		))
	}
}

func (g *Guvnor) statusHandler() {
	g.HTTP.Router.Handle("/status", negroni.New(
		negroni.Wrap(http.HandlerFunc(g.healthCheck)),
	))
}

func (g *Guvnor) healthCheck(w http.ResponseWriter, r *http.Request) {
	response := types.Health{Status: "OK"}
	js, err := json.Marshal(response)
	if err != nil {
		g.Monitoring.Logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
