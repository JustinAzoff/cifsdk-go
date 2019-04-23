package cifsdk

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	resty "gopkg.in/resty.v1"
)

// VERSION
const VERSION = "0.0a3"
const USER_AGENT = "cifsdk-go/" + VERSION

type Indicator struct {
	Id          int      `json:"id"`
	Indicator   string   `json:"indicator"`
	Itype       string   `json:"itype"`
	Portlist    string   `json:"portlist"`
	Firsttime   string   `json:"firsttime"`
	Lasttime    string   `json:"lasttime"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Description string   `json:"description"`
	Count       int      `json:"count"`
	Asn         float32  `json:"asn"`
	Asn_desc    string   `json:"asn_desc"`
	Cc          string   `json:"cc"`
	Provider    string   `json:"provider"`
	Tags        []string `json:"tags"`
	Content     string   `json:"content"`
}

type Feed struct {
	Name        string      `json:"name"`
	User        string      `json:"user"`
	Description string      `json:"description"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
	Indicators  []Indicator `json:"indicators"`
}

func getEnvWithDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

type Client struct {
	Endpoint string
	Token    string
	Debug    bool
}

type indicatorReq struct {
	Indicator *Indicator `json:"indicator"`
}

func (c *Client) CreateIndicator(i *Indicator) error {
	url := fmt.Sprintf("%s/indicators/", c.Endpoint)

	s, err := json.Marshal(&indicatorReq{i})
	if err != nil {
		return err
	}

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", USER_AGENT).
		SetHeader("Authorization", c.Token).
		SetBody(s).
		Post(url)

	if err != nil {
		return err
	}

	if c.Debug {
		spew.Dump(resp)
	}

	return nil
}

func (c *Client) GetFeed(feed string, limit string) (*Feed, error) {
	if limit == "" {
		limit = "25"
	}

	url := fmt.Sprintf("%s/indicators/", c.Endpoint)

	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"limit": limit,
			"itype": feed,
		}).
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", USER_AGENT).
		SetHeader("Authorization", c.Token).
		SetResult(&Feed{}).
		Get(url)

	if err != nil {
		return nil, err
	}
	if c.Debug {
		spew.Dump(resp)
	}

	return resp.Result().(*Feed), nil
}
