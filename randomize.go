package main

import (
  "fmt"
  r "math/rand"
)

func main() {
  /*
  * rand.Intn() is a function from math/rand
  * It takes an integer as the maximum value, and returns a random integer between 0 and the maximum value.
  */
  amount := r.Intn(10000)

  fmt.Println("Amount is: ", amount)

  if amount > 5000 {
    fmt.Println("That's a lot of money.")
  } else {
    fmt.Println("That's not a lot of money.")
  }
}