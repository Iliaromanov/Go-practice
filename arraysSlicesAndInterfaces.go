package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3} // Fixed length of 3
	b := []int{1, 2} // Slice; not fixed length
	c := make([]int, 0, 3) // make creates dynamic array; len=0, cap=3
	// c = c[:3] // re-slice c to make len=3
	c = append(c, 1, 2, 3)
	b = append(b, 3)
	fmt.Printf("a: %v || b: %v || c: %v\n", a, b, c)
	fmt.Println(sliceEquals(b, c))

	greetings := []interface{}{"hi", 1, "hey", true, "ho", 3.14} // Set type as interface to allow any data type in array
	fmt.Printf("This is greetings before removing: %v || ", greetings) // This is greetings before removing: [hi 1 hey true ho 3.14] || The length is: 6
	greetings = remove("hey", greetings)
	fmt.Printf("This is greetings after removing: %v\n", greetings) // This is greetings after removing: [hi 1 true ho 3.14]
	greetings = remove(true, greetings)
	greetings = remove("hi", remove(3.14, remove(true, greetings)))
	fmt.Printf("This is the final version of greetings:\n%v\n", greetings)
}

func remove(target interface{}, arr []interface{}) []interface{} { // Interface gives ability to pass any data type; similar to void* in C
	/* removes target from arr and returns new slice with target removed if found
	   or returns arr if target not found
	*/
	i := 0
	length := len(arr) // cap() can be used to determine the amount of available space in arr
	fmt.Printf("The length is: %v\n", length)
	for i < length && arr[i] != target {
		i++
	}
	if i == length - 1 { // target not found
		return arr
	} 
	new := arr[:i]
	for j := i + 1; j < len(arr); j++ {
		new = append(new, arr[j])
	}
	return new
}

func sliceEquals(nums1 []int, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}