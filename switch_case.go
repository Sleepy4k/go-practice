package main

import "fmt"

func main() {
  name := "John"

  /*
  * Switch is a statement that takes a variable and executes the code block if the variable matches one of the cases.
  */
  switch name {
    case "John":
      fmt.Println("My name is John.")
    case "Doe":
      fmt.Println("My name is Doe.")
    default:
      fmt.Println("My name is unknown.")
  }
}