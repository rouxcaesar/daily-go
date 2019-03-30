// Boiling prints the boiling point of water

package main

import "fmt"

const boilingF = 212.0  // package-level declaration

func main() {
  var f = boilingF            // local to package main
  var c = (f - 32) * 5 / 9    // local to package main
  fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
  // Output:
  // boiling point = 212°F or 100°C
}
