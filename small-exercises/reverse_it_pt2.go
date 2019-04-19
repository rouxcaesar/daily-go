package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseIt("Professional"))          // "lanoisseforP"
	fmt.Println(reverseIt("Walk around the block")) // "Walk dnuora the kcolb"
	fmt.Println(reverseIt("Launch School"))         // "hcnuaL loohcS"
}

func reverseIt(sentence string) string {
	words := strings.Split(sentence, " ")

	for idx, word := range words {
		//fmt.Println("Current word:", word)
		if len(word) >= 5 {
			words[idx] = reverseWord(word)
		}
	}
	//fmt.Println("words after range:", words)
	return strings.Join(words, " ")
}

func reverseWord(word string) string {
	chars := strings.Split(word, "")
	//fmt.Println("reversing word:", word)

	for i := 0; i < len(chars)/2; i++ {
		chars[i], chars[len(chars)-1-i] = chars[len(chars)-1-i], chars[i]
	}
	//fmt.Println("reversed chars:", chars)
	return strings.Join(chars, "")
}
