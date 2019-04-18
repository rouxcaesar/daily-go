// Write a function that counts the number of occurrences of each element in a given array

// Expected Input:
// vehicles = [ 'car', 'car', 'truck', 'car', 'SUV', 'truck', 'motorcycle', 'motorcycle', 'car', 'truck' ]

// Expected Output:
// car => 4
// truck => 3
// SUV => 1
// motorcycle => 2

// Input: a slice of strings that may contain duplicate elements
// Output: print each unique element with a hashrocket and the number of duplicates in the slice

// Approach:
// - Initialize an empty map of string keys to int values
// - iterate through the input slice and for each iteration:
//        - check if element is already a key in the map
//        - if true, increment the value by 1
//        - else, add string key to map with value of 1
// - after iteration is complete, iterate through the map with range
// - for each iteration, fmt.Println(key, "=>", value)

package main

import "fmt"

func main() {
	vehicles := []string{"car", "car", "truck", "car", "SUV", "truck", "motorcycle", "motorcycle", "car", "truck"}

	ledger := howMany(vehicles)
  
  for vehicle, num := range ledger {
    fmt.Println(vehicle, "=>", num)
  }
}

func howMany(vehicles []string) map[string]int {
	ledger := make(map[string]int)

	for _, vehicle := range vehicles {
		_, ok := ledger[vehicle]
		if !ok {
			ledger[vehicle] = 1
		} else {
			ledger[vehicle] += 1
		}
	}

	return ledger
}
