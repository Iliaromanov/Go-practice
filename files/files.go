package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
	"io"
	"log"
)

func main() {
	file, err := os.Open("someText.txt") // For read access
	if err != nil {
		log.Fatal(err)
	}

	// Read file in chunks
	fmt.Println("Chunks:")
	chunkSize := 1
	reader := bufio.NewReader(file)
	buf := make([]byte, chunkSize)

	for {
		n, err := reader.Read(buf) // Read chunk of size n to buf
		
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			} else {
				break
			}
		}
		fmt.Print("|", string(buf[0:n]))
	}
	fmt.Println("\n~~~")
	file.Close()

	file, err = os.Open("someText.txt") // For read access
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // defer waits for rest of function to finish, then this line is executed

	// Read file line by line
	fmt.Println("Line by line:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("-", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("~~~")

	// Read file content all at once
	fmt.Println("All at once:")
	content, err := ioutil.ReadFile("someText.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

/* Stdout:
Chunks:
|H|e|l|l|o|!|
|C|a|n| |y|o|u| |r|e|a|d| |t|h|i|s| |i|n| |G|o|?
~~~
Line by line:
- Hello!
- Can you read this in Go?
~~~
All at once:
Hello!
Can you read this in Go?
*/