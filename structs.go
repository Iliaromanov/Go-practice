package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	// Pointers in Go work the same as pointers in C,
	//  but no pointer arithmetic
	i, j := 5, 6
	var iPtr *int // can declare then assign
	iPtr = &i
	jPtr := &j // can declare and assign in one line
	fmt.Println("i:", *iPtr, "j:", *jPtr) // dereferencing

	// simple struct assignment
	p := Point{1, 2}
	fmt.Println(p, "=> x:", p.x, "y:", p.y)

	// struct assignment with pointers
	p1 := Point{3, 4}
	p1Ptr := &p1
	p1Ptr.x = 5e2 // no need to dereference p1Ptr like in C. (e2 means x10^2)
	fmt.Println(p1, "<=>", *p1Ptr)

	// if field left blank, its initialized to zero
	p2 := Point{x: 1e5} // y=0 implicitly
	p3 := Point{} // x=0 and y=0 implicitly
	fmt.Println(p2, p3)

	// slice of structs
	langs := []struct{
		name string
		known bool
	} {
		{"Go", true},
		{"Python", true},
		{"C", true},
		{"JavaScript", true},
		{"Java", false},
	}

	for _,val := range langs { // need to put _ to avoid "declared but not used" error
		if val.known {
			fmt.Println(val)
		}
	}
}