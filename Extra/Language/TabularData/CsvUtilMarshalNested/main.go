package main

import (
	"fmt"

	"github.com/jszwec/csvutil"
)

func main() {
	type Address struct {
		Street string `csv:"street"`
		City   string `csv:"city"`
	}

	type User struct {
		Name        string  `csv:"name"`
		Address     Address `csv:",inline"`
		HomeAddress Address `csv:"home_address_,inline"`
		WorkAddress Address `csv:"work_address_,inline"`
		Age         int     `csv:"age,omitempty"`
	}

	users := []User{
		{
			Name:        "John",
			Address:     Address{"Washington", "Boston"},
			HomeAddress: Address{"Boylston", "Boston"},
			WorkAddress: Address{"River St", "Cambridge"},
			Age:         26,
		},
	}

	b, err := csvutil.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%s\n", b)

	// Output:
	// name,street,city,home_address_street,home_address_city,work_address_street,work_address_city,age
	// John,Washington,Boston,Boylston,Boston,River St,Cambridge,26
}
