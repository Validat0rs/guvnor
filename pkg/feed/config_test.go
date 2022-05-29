package feed

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"

	"github.com/jaswdr/faker"
)

type configSuite struct {
	feed  *Feed
	faker faker.Faker
}

type readerTest struct {
	fileName string
}

var (
	_       = Suite(&configSuite{})
	chainId = "test-chain-1"
	api     = "https://some-api.url"
)

func Test(t *testing.T) { TestingT(t) }

func (s *configSuite) SetUpSuite(c *C) {
	s.feed = &Feed{}
	s.faker = faker.New()
}

func (s *configSuite) TestParseConfig(c *C) {
	config, err := s.feed.ParseConfig(&readerTest{fileName: s.faker.File().FilenameWithExtension()})
	if err != nil {
		c.Error(err)
	}

	if len(config.Chain) == 0 {
		c.Error("config is empty")
	} else {
		if config.Chain[0].ChainId != chainId {
			c.Error("chain_id does not match")
		}
		if config.Chain[0].API != api {
			c.Error("api does not match")
		}
	}
}

func (s *configSuite) TearDownSuite(c *C) {}

func (r *readerTest) readFile() ([]byte, error) {
	config := fmt.Sprintf(`
[[chain]]
chain_id = "%s"
api = "%s"
`, chainId, api)

	return []byte(config), nil
}
