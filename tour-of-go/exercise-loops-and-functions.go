package main 

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	count := 0

	for i := 0; i < 10; i += 1 {
		z -= (z*z - x) / (2*z)
		fmt.Println("z is", z)
		count += 1
	}

	fmt.Println("count is", count)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(3))
}