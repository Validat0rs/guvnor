package feed

import (
	"log"
	"net/http"
	"os"

	"github.com/Validat0rs/guvnor/pkg/feed/handlers/proposals"

	"github.com/urfave/negroni"
)

func (f *Feed) SetHandlers() {
	f.proposals()
}

func (f *Feed) proposals() {
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
