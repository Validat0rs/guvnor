package feed

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Validat0rs/guvnor/pkg/feed/handlers/proposals"
	"github.com/Validat0rs/guvnor/pkg/feed/types"

	"github.com/urfave/negroni"
)

func (f *Feed) SetHandlers() {
	f.proposalsHandler()

	f.HTTP.Router.Handle("/status", negroni.New(
		negroni.Wrap(http.HandlerFunc(f.healthCheck)),
	))
}

func (f *Feed) proposalsHandler() {
	config, err := f.ParseConfig(&Reader{fileName: os.Getenv("GUVNOR_CONFIG")})
	if err != nil {
		log.Fatal(err)
	}

	for _, chain := range config.Chain {
		_proposals := proposals.NewProposals(chain.ChainId, chain.API, f.Cache.Redis, f.Monitoring.Logger)
		f.HTTP.Router.Handle("/"+chain.ChainId+"/proposals/rss", negroni.New(
			negroni.Wrap(http.HandlerFunc(_proposals.Rss)),
		))
	}
}

func (f *Feed) healthCheck(w http.ResponseWriter, r *http.Request) {
	response := types.Health{Status: "OK"}
	js, err := json.Marshal(response)
	if err != nil {
		f.Monitoring.Logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
