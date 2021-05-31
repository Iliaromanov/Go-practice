package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	//"path/filepath"
)

type Stocks struct {
	Stocks []Stock `json:"stocks"`
}

type Stock struct {
	Name 	    string  `json:"name"`
	ShareValue  float64 `json:"shareValue"`
	LastUpdate  string  `json:"lastUpdate"`
}

func main() {
	f, err := os.Open("stockInfo.json")
	checkError(err)

	// read all at once
	byteValue, err := ioutil.ReadAll(f)
	checkError(err)

	// unmarshal json data into structs defined above
	var stocks Stocks
	json.Unmarshal(byteValue, &stocks)
	println("here:", stocks.Stocks)

	for _, stock := range(stocks.Stocks) {
		fmt.Printf("~~~\nname: %v | value: %v | lastUpdated: %v\n~~~\n", 
			stock.Name, stock.ShareValue, stock.LastUpdate)
	}
}


// Func for error handling
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}