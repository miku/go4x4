package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

type Entry struct {
	XMLName xml.Name  `xml:"entry"`
	Name    string    `xml:"name"`
	Date    time.Time `xml:"t"`
	Refs    []string  `xml:"refs"`
}

func main() {
	entry := Entry{
		Name: "Go",
		Date: time.Now(),
		Refs: []string{"a", "b", "c"},
	}

	// Serialize entry to XML.
	b, err := xml.Marshal(entry)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// Deserialize.
	var v Entry
	err = xml.Unmarshal(b, &v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", v)
}

// $ go run x/encxml/main.go
// <entry><name>Go</name><t>2022-05-03T21:01:30.324596212+02:00</t><refs>a</refs><refs>b</refs><refs>c</refs></entry>
// {XMLName:{Space: Local:entry} Name:Go Date:2022-05-03 21:01:30.324596212 +0200 CEST Refs:[a b c]}
