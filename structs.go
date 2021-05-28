package main

import "fmt"

func main() {
	// Pointers in Go work the same as pointers in C,
	//  but no pointer arithmetic
	i, j := 5, 6
	var iPtr *int // can declare then assign
	iPtr = &i
	jPtr := &j // can declare and assign in one line
	fmt.Println("i:", *iPtr, "j:", *jPtr) // dereferencing
}