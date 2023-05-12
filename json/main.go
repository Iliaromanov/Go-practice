package main

import (
	"fmt"
	"encoding/json"
)

type AllInput struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
	Deps []string `json:"deps"`
}

type OnlyName struct {
	Name string `json:"name"`
}


func main() {
	// type cast string to []byte
	// backticks - `` - are write a string literal; means we dont need to escape special chars like " or {
	sample_json := []byte(`{"name": "COF", "value": 50.5, "deps": ["META", "GOOG", "AMZN"]}`);
	
	// read whole json
	var resultAll AllInput;
	err := json.Unmarshal(sample_json, &resultAll); // parse json into result object
	if err != nil {
		panic(err);
	}

	fmt.Println("resultAll:", resultAll)
	fmt.Printf( // accessing specific fields
		"resultAll.name: %v, resultAll.value: %v, resultAll.deps: %v\n", 
		resultAll.Name, resultAll.Value, resultAll.Deps,
	);

	// read only specific field
	var resultName OnlyName;
	err = json.Unmarshal(sample_json, &resultName); // parse json into result object
	if err != nil {
		panic(err);
	}

	fmt.Println("resultName:", resultName)
	fmt.Printf("resultName.Name: %v\n", resultName.Name)
}