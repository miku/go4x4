// $ go run tasks/btcprice/main.go
// 35995.1852

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CurrentPrice struct {
	Bpi struct {
		EUR struct {
			Code        string  `json:"code"`
			Description string  `json:"description"`
			Rate        string  `json:"rate"`
			RateFloat   float64 `json:"rate_float"`
			Symbol      string  `json:"symbol"`
		}
		GBP struct {
			Code        string  `json:"code"`
			Description string  `json:"description"`
			Rate        string  `json:"rate"`
			RateFloat   float64 `json:"rate_float"`
			Symbol      string  `json:"symbol"`
		}
		USD struct {
			Code        string  `json:"code"`
			Description string  `json:"description"`
			Rate        string  `json:"rate"`
			RateFloat   float64 `json:"rate_float"`
			Symbol      string  `json:"symbol"`
		}
	} `json:"bpi"`
	ChartName  string `json:"chartName"`
	Disclaimer string `json:"disclaimer"`
	Time       struct {
		Updated    string `json:"updated"`
		UpdatedISO string `json:"updatedISO"`
		Updateduk  string `json:"updateduk"`
	} `json:"time"`
}

func main() {
	resp, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	var price CurrentPrice
	if err := dec.Decode(&price); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", price.Bpi.EUR.RateFloat)
}
