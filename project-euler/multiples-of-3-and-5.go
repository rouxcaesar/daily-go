// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

// Input: max value (1000)
// Output: sum of natural numbers below input value that are multiples of 3 or 5

// Natural numbers are positive integers (whole numbers).
// No decimals or negative numbers.

// Approach:
// - Initialize a sum with value 0
// - Use a for loop to iterate from 1 up to input value (exclusive)
// - With each iteration, we'll check if the number is a multiple of 3 or 5.
//        - If true, we'll add it to the sum.
// - At the end of the function, we'll return the sum.

package main 

import "fmt"

// func main() {
// 	sum := 0

// 	for i := 1; i < 1000; i++ {
// 		if (i % 3 == 0) || (i % 5 == 0) {
// 			sum += i
// 		}
// 	}
// 	fmt.Println(sum)
// }

func sumOfMultiples(num int) int {
	sum := 0
	for i := 1; i < num; i++ {
		if isMultiple(i) {
			sum += i
		}
	}
	return sum
}

func isMultiple(num int) bool {
	return (num % 3 == 0) || (num % 5 == 0)
}

func main() {
	fmt.Println("Sum of multiples for input 10 is", sumOfMultiples(10))
	fmt.Println("Sum of multiples for input 1000 is", sumOfMultiples(1000))
}



