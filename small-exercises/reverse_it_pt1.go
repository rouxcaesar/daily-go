// Write a method that takes one argument, a string, and returns a new string with the words in reverse order.
//
// Examples:
//
// puts reverse_sentence('') == ''
// puts reverse_sentence('Hello World') == 'World Hello'
// puts reverse_sentence('Reverse these words') == 'words these Reverse'

// Input: string
// Output: new string with words in reverse order

// DS: slice to represent space-delimited words

// Approach:
// - Take the input string and split by whitespace into words. Store in a slice.
// - Create a new string variable to store reversed words.
// - Iterate through slice of words and concatenate each word to the string variable with a trailing whitespace (except for the last word)

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseString("Hello World"))         // "World Hello"
	fmt.Println(reverseString(""))                    // ""
	fmt.Println(reverseString("Reverse these words")) // "World Hello"

}

func reverseString(sentence string) string {
	words := strings.Split(sentence, " ")

	for i := 0; i < (len(words))/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}

	return strings.Join(words, " ")
}
