package proposals

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Validat0rs/guvnor/pkg/feed/handlers/proposals/types"

	"github.com/gorilla/feeds"
)

var (
	feed = &feeds.Feed{
		Link: &feeds.Link{
			Href: "",
		},
		Description: "Proposals",
		Author: &feeds.Author{
			Name:  os.Getenv("FEED_AUTHOR_NAME"),
			Email: os.Getenv("FEED_AUTHOR_EMAIL"),
		},
		Created: time.Now(),
	}
)

const (
	proposalsUri      = "cosmos/gov/v1beta1/proposals"
	paginationReverse = "pagination.reverse=true"
)

func (p *Proposals) getRawFeed() (*[]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s?%s", p.api, proposalsUri, paginationReverse))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func (p *Proposals) rawFeedToRss(url string, list types.List) (*string, error) {
	feed.Title = p.chainId
	feed.Link.Href = url

	var feedItems []*feeds.Item
	for _, item := range list.Proposals {
		feedItems = append(feedItems,
			&feeds.Item{
				Id:    item.ProposalID,
				Title: item.Content.Title,
				Link: &feeds.Link{
					Href: fmt.Sprintf("%s/%s/%s", p.api, proposalsUri, item.ProposalID),
				},
				Description: item.Content.Description,
				Created:     item.SubmitTime,
			})
	}

	feed.Items = feedItems
	rss, err := feed.ToRss()
	if err != nil {
		return nil, err
	}

	if err := p.setCache(rss); err != nil {
		return nil, err
	}

	return &rss, nil
}
