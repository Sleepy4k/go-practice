package main

import "fmt"

func main() {
  var name string
  var age int
  var height float32

  /*
  * Scan is a function that takes any number of arguments of any type.
  * It takes the values and converts them to the type of the arguments.
  * It returns the number of arguments successfully scanned and an error.
  * Scan is a variadic function, which means it can take any number of arguments.
  */
  fmt.Print("What is your name? ")
  fmt.Scan(&name)
  fmt.Print("")
  fmt.Print("What is your age? ")
  fmt.Scan(&age)
  fmt.Print("")
  fmt.Print("What is your height in meters? ")
  fmt.Scan(&height)
  fmt.Print("")

  fmt.Printf("Your name is %v, you are %v years old and you are %v meters tall.\n", name, age, height)
}