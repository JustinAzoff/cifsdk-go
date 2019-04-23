package cifsdk

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	resty "gopkg.in/resty.v1"
)

// VERSION
const VERSION = "0.0a3"

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

func CreateIndicator(token string, user string, feed string, i *Indicator) bool {
	url := fmt.Sprintf("https://csirtg.io/api/users/%s/feeds/%s/indicators", user, feed)

	s, err := json.Marshal(i)

	var s1 strings.Builder
	s1.WriteString(`{"indicator": `)
	s1.WriteString(string(s))
	s1.WriteString(`}`)

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", fmt.Sprintf("csirtgsdk-go/%s", VERSION)).
		SetAuthToken(token).
		SetBody(s1.String()).
		Post(url)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return false
	}

	debug := getEnvWithDefault("DEBUG", "0")
	if debug == "1" {
		spew.Dump(resp)
	}

	return true
}

func GetFeed(token string, user string, feed string, limit string) *Feed {

	if limit == "" {
		limit = "25"
	}

	url := fmt.Sprintf("https://csirtg.io/api/users/%s/feeds/%s", user, feed)

	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"limit": limit,
		}).
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", fmt.Sprintf("csirtgsdk-go/%s", VERSION)).
		SetAuthToken(token).
		SetResult(&Feed{}).
		Get(url)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	return resp.Result().(*Feed)
}
