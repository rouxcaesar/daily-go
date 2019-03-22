// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

// Input: two nums type int: 999 and 999 -> consider every number from 999 down to 100
// Output: num type int: largest palindrome made from product of two inputs

// Approach:
// - initialize a variable called largestPalindromeProduct and assign it to 0
// - consider the products of two numbers from 999 down to 100 -> (999 * 999), (999 * 998), (999 * 997), etc.
// - check whether the product is a palindrome -> If true, assign value to largestPalindromeProduct
// - return largestPalindromeProduct

package main 

import ( 
	"fmt"
	"strings"
	"strconv"
)

func isPalindromeNum(num int) bool {
	stringNum := strconv.Itoa(num)
	if stringNum == reverseString(stringNum) {
		return true
	}
	return false
}

func reverseString(str string) string {
	var reversed string
	var chars = strings.Split(str, "")
	
	for i := len(chars) - 1; i >= 0; i-- {
		reversed += chars[i]
	}

	return reversed
}

func main() {
	var largestPalindromeProduct int

	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			product := i * j

			if isPalindromeNum(product) && product > largestPalindromeProduct {
				largestPalindromeProduct = product
			}
		}
	}
	fmt.Println(largestPalindromeProduct)	// 906609
}







