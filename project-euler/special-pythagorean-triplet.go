// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a^2 + b^2 = c^2

// For example: 3^2 + 4^2 = 9 + 16 = 25 = 5^2

// There exists exactly one pythagorean triplet for which a + b + c = 1000
// Find the product abc.

// Find the values of a, b, c that satisfy the following requirements:
// 1) a + b + c = 1000
// 2) a < b < c
// 3) a^2 + b^2 = c^2

package main

import "fmt"

func test_triplet(a int, b int, c int) bool {
  return (a*a + b*b) == c*c
}

func main() {
  a, b, c := 0, 0, 0
  for a = 1; a < 1000; a++ {
    for b = 1; b < 1000; b++ {
      for c = 1; c < 1000; c++ {
        if test_triplet(a, b, c) {
          fmt.Printf("Found triplet: %d, %d, %d (sum=%d)\n", a, b, c, a+b+c)
          if a+b+c == 1000 {
            fmt.Printf("%d^2 + %d^2 = %d^2 (%d + %d = %d)\n", a, b, c, a*a, b*b, c*c)
            fmt.Printf("%d x %d x %d = %d\n", a, b, c, a*b*c)
            return
          }
        }
      }
    }
  }
}
