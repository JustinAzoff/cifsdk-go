package cifsdk

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
)

func ToCsv(f *Feed, ow io.Writer) {
	w := csv.NewWriter(ow)

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

	// Write any buffered data to the underlying writer
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
