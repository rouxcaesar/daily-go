// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

// What is the sum of the digits of the number 2^1000?

// Approach:
// 1) Calculate the result of 2^1000 and assign it to a variable
// 2) Initialize a sum variable and assign it to zero
// 3) Create a for loop that will continue until the calculated result is <= 0
// 4) For each iteration, find the result of the calculation % 10 and add it to sum
// 5) Once the loop is done, return sum.

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	calculation := new(big.Int)
	calculation.Exp(big.NewInt(2), big.NewInt(1000), nil)

	//calculation = calculation.Exp(big.NewInt(2), big.NewInt(15), nil)

	s := calculation.String()
	sum := uint64(0)
	for _, c := range s {
		n, _ := strconv.Atoi(string(c))
		sum += uint64(n)
	}

	fmt.Println("Expect sum = 26, we get:", sum)
}
