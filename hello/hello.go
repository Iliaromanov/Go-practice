package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)
import "rsc.io/quote"

func main() {
	log.SetPrefix("greetings: ") // Prints this before any other log messages
	log.SetFlags(0)

	messageFormal, err1 := greetings.HelloFromal("Ilia", "")
	// If error returned print it to console and exit program
	if err1 != nil {
		log.Fatal(err1) // .Fatal prints given argument and stops the program.
	}

	fmt.Println(messageFormal)
	messageInformal, err2 := greetings.HelloInformal("Ilia")
	// If error returned print it to console and exit program
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(messageInformal)
	fmt.Println(quote.Go())
}