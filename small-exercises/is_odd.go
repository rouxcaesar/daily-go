package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isOdd(2))   // false
	fmt.Println(isOdd(5))   // true
	fmt.Println(isOdd(-17)) // true
	fmt.Println(isOdd(-8))  // false
	fmt.Println(isOdd(0))   // false
	fmt.Println(isOdd(7))   // true
}

func isOdd(num int) bool {
	absNum := int(math.Abs(float64(num)))
	return absNum%2 != 0
}
