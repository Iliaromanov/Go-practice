package greetings

import (
	"fmt"
	"errors"
)

/*
*** In Go, a function whose name starts with a capital letter can be called by a function not in the same package. ***
*** This is known as an exported name. ***
*/

func checkEmpty(name string) bool {
	return name == ""
}

// Hello() returns a greeting for the named person.
func HelloFromal(firstName string, lastName string) (string, error) { // has to be capital to be used in other file

	if checkEmpty(firstName) || checkEmpty(lastName) {
		return "", errors.New("empty name field")
	}

	// Return a greeting that embeds the name in a message.
	msg := fmt.Sprintf("Hello, %v %v. Welcome!", firstName, lastName) // Create a formatted string

	return msg, nil
}

func HelloInformal(name string) (string, error) {

	if checkEmpty(name) {
		return "", errors.New("empty name field")
	}
	var msg string
	msg = fmt.Sprintf("Hey %v!", name)
	msg += " What's up???"
	return msg, nil
}