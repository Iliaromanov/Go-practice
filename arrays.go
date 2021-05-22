package main

import (
	"fmt"
)

// Interface gives ability to pass any data type; similar to void* in C
func remove(target interface{}, arr []interface{}) []interface {
	//i := 0
	//for arr[i] != target {
	//	i++
	//}
}

func main() {
	a := [3]int{1, 2, 3} // Fixed length of 3
	b := []int{1, 2} // Slice; not fixed length
	b = append(b, 3)
	fmt.Println(a, b)

	greetings := []string{"hi", "hey", "ho"}
	greetings = append(greetings[:1], )

}