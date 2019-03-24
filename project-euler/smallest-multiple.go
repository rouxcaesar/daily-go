// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

// Notes: 2520 is the first number encountered (starting at 1 and incrementing) that whose modulo equals zero for all numbers from 1 upto 10 (inclusive)

// Brute Force: 
// - Starting from 21, we will increment by 1 until we find our target number.
// - Our target number will need to evaluate to true for (target % num == 0) where num is each value in the range (1..20)

package main 

import "fmt"

func isTarget(num int) bool {
	for i := 1; i <= 20; i++ {
		if num % i != 0 {
			return false
		}
	}

	return true
}

func main() {
	for i := 21; i > 0; i++ {
		if isTarget(i) {
			fmt.Println(i)
			return
		}
	}
}