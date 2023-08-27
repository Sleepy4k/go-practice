package main

import (
  "fmt"
  t "time"
  r "math/rand"
)

func main() {
  /*
  * rand.Seed() is a function from math/rand
  * It takes an integer as the seed value, and returns a random integer between 0 and the maximum value.
  * The seed value is usually the current time in nanoseconds.
  * This ensures that the random number is different every time the program is run.
  *
  * time.Now() is a function from time
  * It returns the current time.
  * time.Now().UnixNano() returns the current time in nanoseconds.
  *
  * Btw this method is not recommended to use in production or desprecated in Go 1.10
  * https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
  * use randomize method instead
  */
  r.Seed(t.Now().UnixNano())

  amount := r.Intn(10000)

  fmt.Println("Amount is: ", amount)

  if amount > 5000 {
    fmt.Println("That's a lot of money.")
  } else {
    fmt.Println("That's not a lot of money.")
  }
}