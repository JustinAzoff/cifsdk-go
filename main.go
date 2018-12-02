package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"gopkg.in/resty.v1"
	"log"
	"os"
)

// VERSION
const VERSION = "0.0a0"

type Feed struct {
	Name string `json:"name"`
	User string `json:"user"`
	Description string `json:"description"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Indicators []struct {
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
		Asn string `json:"asn"`
		Asn_desc string `json:"asn_desc"`
		Cc string `json:"cc"`
		Provider string `json:"provider"`
		Tags []string `json:"tags"`
		Content string `json:"content"`
	} `json:"indicators"`
}



//https://www.scaledrone.com/blog/creating-an-api-client-in-go/

func main() {

	user := flag.String("user", "csirtgadgets", "user")
	feed := flag.String( "feed", "darknet", "feed name" )
	token := os.Getenv("CSIRTG_TOKEN")

	flag.Parse()

	url := fmt.Sprintf("https://csirtg.io/api/users/%s/feeds/%s", *user, *feed)

	resty.SetDebug(true)
	resp, err := resty.R().
		SetQueryParams(map[string]string{
		"limit": "20",
		}).
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", fmt.Sprintf("csirtgsdk-go/%s", VERSION)).
		SetAuthToken(token).
		SetResult(&Feed{}).
		Get(url)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	var f = resp.Result().(*Feed)
	//spew.Dump(f.Indicators)

	w := csv.NewWriter(os.Stdout)

	for _, i := range f.Indicators {
		r := []string{
			i.Itype,
			i.Indicator,
			i.Asn,
			i.Asn_desc,
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
