package main

import (
	"flag"
	"fmt"
	"gopkg.in/resty.v1"
	"log"
	"os"
)

// VERSION
const VERSION = "0.0a0"


func main() {

	user := flag.String("user", "csirtgadgets", "user")
	feed := flag.String( "feed", "darknet", "feed name" )
	token := os.Getenv("CSIRTG_TOKEN")

	flag.Parse()

	url := fmt.Sprintf("https://csirtg.io/api/users/%s/feeds/%s.csv", *user, *feed)

	resp, err := resty.R().
		SetQueryParams(map[string]string{
		"limit": "20",
		}).
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", fmt.Sprintf("csirtgsdk-go/%s", VERSION)).
		SetAuthToken(token).
		Get(url)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	fmt.Println(resp)
}
