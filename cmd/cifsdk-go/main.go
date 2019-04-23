package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	cif "github.com/JustinAzoff/cifsdk-go"
	"github.com/davecgh/go-spew/spew"
)

//https://www.scaledrone.com/blog/creating-an-api-client-in-go/

func main() {

	endpoint := flag.String("endpoint", "http://localhost:5000", "cif endpoint")
	feed := flag.String("feed", "darknet", "feed name")
	limit := flag.String("limit", "25", "result limit")
	format := flag.String("format", "csv", "output format")
	debug := flag.Bool("debug", false, "turn on debugging")
	token := os.Getenv("CIF_TOKEN")

	indicator := flag.String("indicator", "", "set indicator")
	tags := flag.String("tags", "", "ssh,scanner,...")
	description := flag.String("description", "", "honeypot scanner")

	flag.Parse()

	//if *debug == true {
	//	resty.SetDebug(true)
	//}

	c := &cif.Client{
		Endpoint: *endpoint,
		Token:    token,
		Debug:    *debug,
	}

	if *indicator != "" {
		var i = &cif.Indicator{
			Indicator:   *indicator,
			Tags:        strings.Split(*tags, ","),
			Description: *description,
		}

		var r = c.CreateIndicator(i)
		spew.Dump(r)
	} else {
		var f, err = c.GetFeed(*feed, *limit)
		if err != nil {
			log.Fatal(err)
		}

		if *format == "csv" {
			cif.ToCsv(f, os.Stdout)
		} else {
			fmt.Println("Format doesn't exist yet, SEND US A PR!")
		}
	}

}
