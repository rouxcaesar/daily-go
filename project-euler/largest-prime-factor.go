// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

// Input: whole number
// Output: whole number that represents the largest prime factor of the input

// Sub-problems:
// - Determine whether a number is prime 				-> only factors are itself and 1
// - Determine whether a number is a factor 		-> input % num == 0
// - Do we work our way bottom-up or top-down?	-> top-down b/c the first prime factor we encounter will be the largest

// Approach:
// - Initialize a variable i and assign to (input - 1)
// - Using a for loop, check each value from i down to 2 and check if the current value of i is a prime factor.
// - Create two helper methods: one to check if a number is prime & one to check if a number is a factor (Have this second one determine if we check if prime -> optimization)
// - Return the first value that returns true for both helper methods.

package main 

import (
	"fmt"
	"math"
)

func sqrt(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}

func isPrime(num int64) bool {
	if num <= 1 { return false }

	for i := sqrt(num); i >= 1; i-- {
		if i == 1 { return true }
		if num % i == 0 { return false }
	}
	return true
}

func isFactor(num, input int64) bool {
	return input % num == 0
}

func findLargestPrimeFactor(num int64) int64 {
	var i int64

	for i = sqrt(num); i >= 1; i-- {
		if isFactor(i, num) && isPrime(i) {
			return i
		}
	}
	return num
}

func main() {
	// fmt.Println(findLargestPrimeFactor(13195)) // 29
	fmt.Println(findLargestPrimeFactor(600851475143)) // 6857
}

