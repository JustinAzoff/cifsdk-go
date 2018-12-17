package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/resty.v1"
	"log"
	"os"
	"strconv"
	"strings"
)

// VERSION
const VERSION = "0.0a2"

type Indicator struct {
	Id int `json:"id"`
	Indicator string `json:"indicator"`
	Itype string `json:"itype"`
	Portlist string `json:"portlist"`
	Firsttime string `json:"firsttime"`
	Lasttime string `json:"lasttime"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Description string `json:"description"`
	Count int `json:"count"`
	Asn float32 `json:"asn"`
	Asn_desc string `json:"asn_desc"`
	Cc string `json:"cc"`
	Provider string `json:"provider"`
	Tags []string `json:"tags"`
	Content string `json:"content"`
}

type Feed struct {
	Name string `json:"name"`
	User string `json:"user"`
	Description string `json:"description"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Indicators []Indicator `json:"indicators"`
}

func toCsv(f *Feed) {

	w := csv.NewWriter(os.Stdout)

	for _, i := range f.Indicators {
		r := []string{
			//strconv.Itoa(i.Id),
			i.Indicator,
			i.Itype,
			i.Portlist,
			i.Firsttime,
			strconv.Itoa(i.Count),
			fmt.Sprintf("%0.f", i.Asn),
			i.Asn_desc,
			i.Description,
			i.Provider,
		}

		if err := w.Write(r); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}


func createIndicator(token string, user string, feed string, i *Indicator) bool {
	url := fmt.Sprintf("https://csirtg.io/api/users/%s/feeds/%s/indicators", user, feed)

	fmt.Println(url)

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

	spew.Dump(resp)

	return true
}

func getFeed(token string, user string, feed string, limit string) *Feed {

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



//https://www.scaledrone.com/blog/creating-an-api-client-in-go/

func main() {

	user := flag.String("user", "csirtgadgets", "user")
	feed := flag.String( "feed", "darknet", "feed name" )
	limit := flag.String("limit", "25", "result limit")
	format := flag.String("format", "csv", "output format")
	debug := flag.Bool("debug", false, "turn on debugging")
	token := os.Getenv("CSIRTG_TOKEN")

	indicator := flag.String("indicator", "", "set indicator" )
	tags := flag.String("tags", "", "ssh,scanner,...")
	description := flag.String("description", "", "honeypot scanner")

	flag.Parse()

	if *debug == true {
		resty.SetDebug(true)
	}

	if *indicator != "" {
		var i = &Indicator{
			Indicator: *indicator,
			Tags: strings.Split(*tags, ","),
			Description: *description,
		}

		var r = createIndicator(token, *user, *feed, i)
		spew.Dump(r)
	} else {
		var f = getFeed(token, *user, *feed, *limit)

		if *format == "csv" {
			toCsv(f)
		} else {
			fmt.Println("Format doesn't exist yet, SEND US A PR!")
		}
	}






}
