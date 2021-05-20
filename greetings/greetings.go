package greetings

import "fmt"

/*
*** In Go, a function whose name starts with a capital letter can be called by a function not in the same package. ***
*** This is known as an exported name. ***
*/

// Hello returns a greeting for the named person.
func HelloFromal(firstName string, lastName string) string { 
	// Return a greeting that embeds the name in a message.
	msg := fmt.Sprintf("Hello, %v %v. Welcome!", firstName, lastName) // Create a formatted string
	return msg
}

func HelloInformal(name string) string {
	var msg string
	msg = fmt.Sprintf("Hey %v!", name)
	msg += " What's up???"
	return msg
}