package main 

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i 						// generates a pointer to i
	fmt.Println(*p)			// read i through pointer (42)

	*p = 21							// set i through pointer
	fmt.Println(i) 			// see the new value of i (21)

	p = &j
	*p = *p / 37				// divide j through the pointer
	fmt.Println(j)			// see the new value of j (73)
}