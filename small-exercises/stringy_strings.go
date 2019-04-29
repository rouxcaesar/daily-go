// Write a method that takes one argument, a positive integer.
// It should return a string of alternating 1s and 0s, always starting with 1.
// The length of the string should match the given integer.

// Input: int
// Output: string

// Algorithm:
//  - Initialize an empty string variable
//  - Start a loop from i = 1 and end when i > input
//  - For each iteration of loop, concat either a '1' or '0' onto the string variable
//      - If current value of i is odd, concat a '1'
//      - Else, concat a '0'
//  - Once loop is done, return the string

package main

import "fmt"

func main() {
	fmt.Println(stringy(6)) // '101010'
	fmt.Println(stringy(9)) // '101010101'
	fmt.Println(stringy(4)) // '1010'
	fmt.Println(stringy(7)) // '1010101'
}

func stringy(num int) string {
  result := ""	

	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			result += "1" 
		} else {
			result += "0" 
		}
	} 
	return result
}
