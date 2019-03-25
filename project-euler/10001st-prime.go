// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.

// What is the 10,001st prime number?

// Approach:
// - initialize a counter and assign to 0
// - start incrementing up from 1 with a for loop
// - have the loop stop once counter == 10001
// - for each current value in the loop, check if it is prime
// - if true, increment counter
// - once the counter == 10001, return the current value as our target prime

package main 

import "fmt"

func isPrime(num int) bool {

	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}

	return true
}

func main() {
	counter := 0

	for i := 2; counter <= 10001; i++ {
		if isPrime(i) {
			counter++
		}

		if counter == 10001 {
			fmt.Println(i)
			break
		}
	}
}