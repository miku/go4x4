package main

import (
	"encoding/csv"
	"io"
	"log"
	"strings"

	"github.com/jszwec/csvutil"
)

func main() {
	r := strings.NewReader(`
a,b
1,2
3,4
NA,6`)
	dec, err := csvutil.NewDecoder(csv.NewReader(r))
	if err != nil {
		log.Fatal(err)
	}

	dec.Map = func(field, column string, v interface{}) string {
		if field == "NA" {
			return "0"
		}
		return field
	}

	// note: only public fields will be decoded
	var data struct {
		A float64 `csv:"a"`
		B string  `csv:"b"`
	}

	for {
		if err := dec.Decode(&data); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		log.Printf("%+v", data)
	}

}
