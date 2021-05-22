package main

import "fmt"

func printMap(m map[string]interface{}) {
	for key, val := range m {
		fmt.Println(key,":", val)
	}
}

func main() {
	balances := make(map[string]interface{})
	balances["Bob"] = 0.25
	balances["Alice"] = 5.50
	balances["Bartholamew"] = 1500.11
	printMap(balances) // >> map[Alice:5.5 Bartholamew:1500.11 Bob:0.25]

	var langs = map[string]interface{} {
		"Python": "$$$$$",
		"Go": "$$$$$",
		"C": "$$$",
		"SQL": "$",
	}
	langs["JavaScript"] = "$$"
	val, ok := langs["SQL"]
	fmt.Printf("val: %v, ok: %t\n", val, ok) // >> val: $, ok: true
	delete(langs, "SQL") // use 'delete' to remove key-value pairs
	val, ok = langs["SQL"]
	fmt.Printf("val: %v, ok: %t\n", val, ok) // >> val: , ok: false
	printMap(langs)// >> map[C:$$$ Go:$$$$$ JavaScript:$$ Python:$$$$$]
}