package main

import "fmt"

func main() {
  var array1 [5]string
  array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

  array1 = array2

  fmt.Println("array1", array1)
  fmt.Println("array2", array2)

  array1[0] = "WHITE"

  fmt.Println("array1", array1)
  fmt.Println("array2", array2)
}


