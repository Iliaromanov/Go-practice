package main

import (
	"fmt"
	// Documentation on fmt package: https://golang.org/pkg/fmt/
)

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

func main() {
	fmt.Println("This is the start of the program!")
	countFromTo(0, 5)
	sevenOdd := isOdd(7) // Assigning to variable
	fmt.Printf("7 is odd? - %t\n", sevenOdd) // Can pass variable
	fmt.Printf("8 is odd? - %t\n", isOdd(8)) // Can pass func result directly
}