package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  fmt.Println("Command invoked:", os.Args[0]) // Exercise 1.1

  // Exercise 1.2
  for idx, arg := range os.Args[1:] {
    fmt.Println("Index:", idx)
    fmt.Println("Arg:", arg)
  }

  fmt.Println(strings.Join(os.Args[1:], " "))
}
