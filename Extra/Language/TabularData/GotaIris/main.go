package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	df := dataframe.ReadCSV(os.Stdin,
		dataframe.DefaultType(series.String),
		dataframe.DetectTypes(true),
		dataframe.HasHeader(false),
		dataframe.Names("sl", "sw", "pl", "pw", "species"),
		dataframe.NaNValues([]string{"NA", "NaN", "<nil>"}))

	fmt.Println(df)

	// Get a column.
	sl := df.Col("sl")

	// Basic stats.
	fmt.Println(sl)

	// Basic stats.
	mean := func(s series.Series) series.Series {
		floats := s.Float()
		sum := 0.0
		for _, f := range floats {
			sum += f
		}
		return series.Floats(sum / float64(len(floats)))
	}

	// Apply, will not err on non-numeric columns.
	cmean := df.Capply(mean)
	rmean := df.Select([]string{"sl", "sw", "pl", "pw"}).Rapply(mean)

	fmt.Println(cmean)
	fmt.Println(rmean)

	fi := df.Filter(dataframe.F{1, "sw", ">", 2.0}).
		Filter(dataframe.F{4, "species", series.Eq, "Iris-setosa"}).
		Select([]string{"pw", "pl"}).
		Subset([]int{0, 2})

	if fi.Err != nil {
		log.Fatal(fi.Err)
	}
	// Overview.
	fmt.Println("describe -->")
	fmt.Println(df.Describe())

	// Filter.
	fmt.Println("filtered -->")
	fmt.Println(fi)

	// As records.
	fmt.Println("records -->")
	fmt.Println(fi.Records())
}
