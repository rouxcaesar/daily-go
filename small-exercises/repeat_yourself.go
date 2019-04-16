package main

import "fmt"

func main() {
	repeat("Hello", 3)
}

func repeat(s string, num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(s)
	}
}
