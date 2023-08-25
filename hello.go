package main

import "fmt"

/*
* The main function is the entry point of the program.
* When we run the program, it is automatically executed first.
*/
func main() {
  /*
  * For print a string in console, we use fmt.Println
  * fmt is a package that contains functions for formatting text, including printing to the console.
  */
  fmt.Println("Hello World!")

  /*
  * For print a string in console, we use fmt.Print
  * fmt.Print is similar to fmt.Println, except that it does not add a new line after the text.
  */
  fmt.Print("My favorite number is ", 9, "\n")

  /*
  * We can also use fmt.Printf to print a formatted string.
  * For example:
  * fmt.Printf("My favorite number is %d\n", 9)
  *
  * %d is a placeholder for a number.
  * %s is a placeholder for a string.
  * %f is a placeholder for a float.
  * %t is a placeholder for a boolean.
  * %v is a placeholder for a variable.
  */
  weight := 70

  fmt.Printf("Days in a week is %d days\n", 7)
  fmt.Printf("My name is %s\n", "John")
  fmt.Printf("My height is %f\n", 170.5)
  fmt.Printf("Are you learn go? %t\n", true)
  fmt.Printf("My weight is %v\n", weight)
}