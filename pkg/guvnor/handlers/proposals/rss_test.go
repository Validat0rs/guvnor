package proposals

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	. "gopkg.in/check.v1"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Validat0rs/guvnor/pkg/guvnor/handlers/proposals/types"

	"github.com/go-redis/redismock/v8"
	"github.com/rs/zerolog/log"
)

type rssSuite struct {
	proposal   *Proposals
	httpServer *httptest.Server
}

type rssItem []struct {
	Text        string `xml:",chardata"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Guid        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
}

type rssChanel struct {
	Text        string  `xml:",chardata"`
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	Description string  `xml:"description"`
	PubDate     string  `xml:"pubDate"`
	Item        rssItem `xml:"item"`
}

type rssResponse struct {
	XMLName xml.Name  `xml:"rss"`
	Text    string    `xml:",chardata"`
	Version string    `xml:"version,attr"`
	Content string    `xml:"content,attr"`
	Channel rssChanel `xml:"channel"`
}

var (
	_       = Suite(&rssSuite{})
	chainId = "test-chain-1"

	proposalItemTitle       = "A new proposal"
	proposalItemDescription = "A really good proposal"
	proposalItemStatus      = "PROPOSAL_STATUS_VOTING_PERIOD"
	proposal                = types.Proposals{{
		ProposalID: "1",
		Content: types.Content{
			Type:        "/cosmos.gov.v1beta1.TextProposal",
			Title:       proposalItemTitle,
			Description: proposalItemDescription,
		},
		Status:         proposalItemStatus,
		SubmitTime:     time.Now(),
		DepositEndTime: time.Now(),
		TotalDeposit: types.TotalDeposit{{
			Denom:  "utoken",
			Amount: "1000000",
		}},
		VotingStartTime: time.Now(),
		VotingEndTime:   time.Now(),
	}}
)

func Test(t *testing.T) { TestingT(t) }

func (s *rssSuite) SetUpSuite(c *C) {
	var list types.List
	list.Proposals = proposal

	js, _ := json.Marshal(list)
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(js)
	})
	httpServer := httptest.NewServer(handlerFunc)

	redisClient, mock := redismock.NewClientMock()
	mock.Regexp().ExpectSet(`[a-z]+`, `[a-z]+`, 6*time.Hour).SetVal("OK")

	s.proposal = NewProposals(chainId,
		fmt.Sprintf("http://%s", httpServer.Listener.Addr().String()),
		redisClient,
		log.With().Str("module", "feed").Logger(),
	)
	s.httpServer = httpServer
}

func (s *rssSuite) TestRss(c *C) {
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", "/", proposalsUri), nil)
	w := httptest.NewRecorder()

	s.proposal.Rss(w, r)
	res := w.Result()
	defer res.Body.Close()

	rss, err := io.ReadAll(res.Body)
	if err != nil {
		c.Error(err)
	}

	var formatted rssResponse
	if err = xml.Unmarshal(rss, &formatted); err != nil {
		c.Error(err)
	}

	if formatted.Channel.Title != chainId {
		c.Error("title does not match")
	}
	if formatted.Channel.Description != "Proposals" {
		c.Error("description does not match")
	}
	if formatted.Channel.Item[0].Title != proposalItemTitle {
		c.Error("item title does not match")
	}
	if formatted.Channel.Item[0].Description != proposalItemDescription {
		c.Error("item description does not match")
	}
}

func (s *rssSuite) TearDownSuite(c *C) {
	s.httpServer.Close()
}
