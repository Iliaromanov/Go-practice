package main

import (
	"fmt"
	// Documentation on fmt package: https://golang.org/pkg/fmt/
)


func main() {
	fmt.Println("This is the start of the program!")
	countFromTo(0, 5)
	sevenOdd := isOdd(7) // Assigning to variable
	fmt.Printf("7 is odd? - %t\n", sevenOdd) // Can pass variable
	fmt.Printf("8 is odd? - %t\n", isOdd(8)) // Can pass func result directly
	fmt.Println(makeBox(4, 5))
}

func countFromTo(start int, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}

func isOdd(num int) bool {
	if num % 2 != 0 {
		return true
	} else {
		return false
	}
		
}

func makeBox(width int, height int) string {
	var box string;
	box = fmt.Sprintf("This is a %v by %v box.", width, height)
	// Two lines above equivalent to "box := fmt.Sprintf("This is a %v by %v box.", width, height)"
	// := determines what type the variable is based on the value thats being assigned

	return box
}