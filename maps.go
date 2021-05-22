package main

import "fmt"

func main() {
	balances := make(map[string]float32)
	balances["Bob"] = 0.25
	balances["Alice"] = 5.50
	balances["Bartholamew"] = 1500.11
	fmt.Print(balances) // >> map[Alice:5.5 Bartholamew:1500.11 Bob:0.25]

	var langs = map[string]string{
		"Python": "$$$$$",
		"Go": "$$$$$",
		"C": "$$$",
		"SQL": "$",
		"JavaScript": "$$",	
	}
	fmt.Print(langs)// >> map[C:$$$ Go:$$$$$ JavaScript:$$ Python:$$$$$ SQL:$]
}