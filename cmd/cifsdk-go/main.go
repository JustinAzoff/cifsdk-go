package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	cif "github.com/JustinAzoff/cifsdk-go"
	"github.com/davecgh/go-spew/spew"
)

//https://www.scaledrone.com/blog/creating-an-api-client-in-go/

func main() {

	user := flag.String("user", "csirtgadgets", "user")
	feed := flag.String("feed", "darknet", "feed name")
	limit := flag.String("limit", "25", "result limit")
	format := flag.String("format", "csv", "output format")
	//debug := flag.Bool("debug", false, "turn on debugging")
	token := os.Getenv("CSIRTG_TOKEN")

	indicator := flag.String("indicator", "", "set indicator")
	tags := flag.String("tags", "", "ssh,scanner,...")
	description := flag.String("description", "", "honeypot scanner")

	flag.Parse()

	//if *debug == true {
	//	resty.SetDebug(true)
	//}

	if *indicator != "" {
		var i = &cif.Indicator{
			Indicator:   *indicator,
			Tags:        strings.Split(*tags, ","),
			Description: *description,
		}

		var r = cif.CreateIndicator(token, *user, *feed, i)
		spew.Dump(r)
	} else {
		var f = cif.GetFeed(token, *user, *feed, *limit)

		if *format == "csv" {
			cif.ToCsv(f, os.Stdout)
		} else {
			fmt.Println("Format doesn't exist yet, SEND US A PR!")
		}
	}

}
