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

	for _, stock := range(stocks.Stocks) {
		fmt.Printf("~~~\nname: %v | value: %v | lastUpdated: %v\n~~~\n", 
			stock.Name, stock.ShareValue, stock.LastUpdate)
	}
	f.Close()

	// Reopen file
	f, err = os.Open("stockInfo.json")
	checkError(err)
	defer f.Close()

	// Reading JSON without using structs (using interface{} instead)
	byteValue2, err := ioutil.ReadAll(f)
	
	checkError(err)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue2), &result)
	fmt.Println(result["stocks"])

}


// Func for error handling
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}