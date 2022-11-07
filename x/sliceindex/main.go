package main

import "log"

func main() {

	// zero indexed, by default
	zMonths := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
	log.Println(len(zMonths), zMonths)
	log.Println(zMonths[0])

	// custom indices possible
	months := []string{
		3:  "mar",
		1:  "jan",
		2:  "feb",
		4:  "apr",
		5:  "may",
		6:  "jun",
		7:  "jul",
		8:  "aug",
		9:  "sep",
		10: "oct",
		11: "nov",
		12: "dec",
	}
	log.Println(len(months), months)
	log.Println(months[1])
	log.Println(months[0]) // does not panic, just returns the empty string

	// log.Println(months[13]) // panic: runtime error: index out of range [13] with length 13

}
