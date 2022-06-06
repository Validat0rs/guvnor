package proposals

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Validat0rs/guvnor/pkg/feed/handlers/proposals/types"
)

func (p *Proposals) Rss(w http.ResponseWriter, r *http.Request) {
	content, err := p.getRawFeed()
	if err != nil {
		p.logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var list types.List

	if err := json.Unmarshal(*content, &list); err != nil {
		p.logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var feed *string
	feed = p.getCache()

	if *feed == "" {
		feed, err = p.rawFeedToRss(fmt.Sprintf("https://%s%s", r.Host, r.URL.String()), list)
		if err != nil {
			p.logger.Error().Msgf("%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/xml")
	_, _ = w.Write([]byte(*feed))
}
