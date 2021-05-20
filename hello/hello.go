package main

import (
	"fmt"
	"example.com/greetings"
)
import "rsc.io/quote"

func main() {
	messageFormal := greetings.HelloFromal("Ilia", "Romanov")
	fmt.Println(messageFormal)
	messageInformal := greetings.HelloInformal("Ilia")
	fmt.Println(messageInformal)
	fmt.Println(quote.Go())
}