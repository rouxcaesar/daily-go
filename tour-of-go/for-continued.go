// The init and post statements of a for loop are optional.
// The condition statement is required.

package main 

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}

	fmt.Println(sum)
}