package main

import "fmt"

func main() {
  isSunday := true

  /*
  * If is a statement that takes a boolean expression and executes the code block if the expression is true.
  */
  if isSunday {
    fmt.Println("Today is Sunday.")
  }

  isSunday = false

  /*
  * If-else is a statement that takes a boolean expression and executes the code block after the else if the expression is false.
  */
  if isSunday {
    fmt.Println("Today is Sunday.")
  } else {
    fmt.Println("Today is not Sunday.")
  }
}