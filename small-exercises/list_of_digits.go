// Write a method that takes one argument, a positive integer, and returns a list of the digits in the number

// Ex: Given 12345, return the array [1,2,3,4,5]

package main

import "fmt"

func main() {
	fmt.Println(digitList(12345))   // expect: [1,2,3,4,5]
  fmt.Println(digitList(7))       // expect: [7]
  fmt.Println(digitList(375290))  // expect: [3, 7, 5, 2, 9, 0]
  fmt.Println(digitList(444))     // expect: [4, 4, 4]
}

func digitList(num int) []int {
	// Initialize an empty slice to store digits
	// Use a loop to construct the slice
	// For each iteration, find the modulo 10 of the input
	// append it to the slice
	// modify the input by dividing by 10 (truncates)
	// Repeat this process until the input <= 0
	// After the loop, return the slice

	digits := make([]int, 0)

	for num > 0 {
		digit := num % 10
		digits = append(digits, digit)
		//fmt.Println("digits:", digits)
		num = num / 10
	}

	digits = reverseSlice(digits)
	return digits
}

func reverseSlice(nums []int) []int {
	length := len(nums)
  //fmt.Println(((len(nums) - 1)/2))
	for i := 0; i < ((len(nums) - 1)/2); i++ {
   // fmt.Println("current index:", i)
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}

	return nums
}
