// The sum of the squares of the first ten natural numbers is,

// 1^2 + 2^2 + ... + 10^2 = 385

// The square of the sum of the first ten natural numbers is,

// (1 + 2 + ... + 10)^2 = 552 = 3025

// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.


// Input: 100 -> represents the first one hundred natural numbers
// Output: the difference between the two calculations -> (sum of squares) - (square of sums)

// Approach:
// - Initialize two variables: sumOfSquares and squareOfSum
// - Construct a for loop that iterates through 1..input. For each iteration, add the square of the current num to sumOfSquares. Also add the current num to squareOfSum.
// - After the loop, return the absolute value of the difference: (sumOfSquares) - (squareOfSum)^2 -> take absolute value


package main

import (
  "fmt"
  "math"
)

func main() {
  sumOfSquares := 0
  squareOfSum := 0

  for i := 1; i <= 100; i++ {
    squareOfSum += i
    sumOfSquares += int(math.Pow(float64(i), 2))
  }

 fmt.Println(math.Abs(float64(sumOfSquares) - (math.Pow(float64(squareOfSum),2))))
 // 25164150
}








