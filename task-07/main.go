package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	//Empty two dimensional string array, because later I am going to append elements into it
	//And append doesnt take single dimensional array
	csvContents := [][]string{}
	geziyor.NewGeziyor(&geziyor.Options{
		//StartRequestsFunc to make request omn my own
		StartRequestsFunc: func(getter *geziyor.Geziyor) {
			getter.GetRendered("https://www.forbes.com/real-time-billionaires/#4502d58a3d78", getter.Opt.ParseFunc)
		},
		//This functions parses the response received from the rquest above
		ParseFunc: func(getter *geziyor.Geziyor, res *client.Response) {
			//Each function gives two parameters,1)index of element 2)the element that satisfies the given condition
			res.HTMLDoc.Find("tr.base.ng-scope").Each(func(index int, element *goquery.Selection) {
				if index < 10 {
					csvContents = append(csvContents, []string{element.Find("td.rank").Text(), element.Find("td.name").Text(), element.Find("td.Net.Worth").Text(), element.Find("td.age").Text(), element.Find("td.Country\\/Territory").Text(), element.Find("td.source").Text()})
				}
			})
		},
	}).Start()
	data, error := os.Create("data.csv")
	if error != nil {
		fmt.Println(error)
	}
	csvWriter := csv.NewWriter(data)
	//blank variable allows us to have an unused variable in GoLang
	for _, m := range csvContents {
		_ = csvWriter.Write(m)
	}
	csvWriter.Flush()
	data.Close()
}
