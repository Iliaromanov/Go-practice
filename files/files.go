package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
	"log"
)

func main() {
	file, err := os.Open("someText.txt") // For read access
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // defer waits for rest of function to finish, then this line is executed

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("-", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Read file content all at once
	content, err := ioutil.ReadFile("someText.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("~~~\n%v", string(content))
}