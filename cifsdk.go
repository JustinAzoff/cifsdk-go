package cifsdk

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

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
type IndicatorList []Indicator

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

func (c *Client) CreateIndicators(i *IndicatorList) error {
	resty.SetTimeout(15 * time.Second)
	url := fmt.Sprintf("%s/indicators/", c.Endpoint)

	s, err := json.Marshal(i)
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

func (c *Client) GetIndicators(itype string, limit string) (IndicatorList, error) {
	resty.SetTimeout(15 * time.Second)
	if limit == "" {
		limit = "25"
	}
	url := fmt.Sprintf("%s/indicators/", c.Endpoint)

	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"limit": limit,
			"itype": itype,
		}).
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", USER_AGENT).
		SetHeader("Authorization", c.Token).
		SetResult(IndicatorList{}).
		Get(url)

	if err != nil {
		return nil, err
	}
	if c.Debug {
		spew.Dump(resp)
	}

	lst := resp.Result().(*IndicatorList)
	return *lst, nil
}
