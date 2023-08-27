package main

import "fmt"

func main() {
  memberCard := "12-34-56-78"
  systemCard := "41-23-65-87"

  /*
  * == is an operator that takes two values of the same type and returns a boolean.
  * It returns true if the values are equal and false if they are not.
  */
  if memberCard == systemCard {
    fmt.Println("Access granted.")
  }

  /*
  * != is an operator that takes two values of the same type and returns a boolean.
  * It returns true if the values are not equal and false if they are.
  */
  if memberCard != systemCard {
    fmt.Println("Access denied.")
  }

  cookies := 5

  /*
  * > is an operator that takes two values of the same type and returns a boolean.
  * It returns true if the first value is greater than the second and false if it is not.
  *
  * >= is an operator that takes two values of the same type and returns a boolean.
  * It returns true if the first value is greater than or equal to the second and false if it is not.
  *
  * < is an operator that takes two values of the same type and returns a boolean.
  * It returns true if the first value is less than the second and false if it is not.
  *
  * <= is an operator that takes two values of the same type and returns a boolean.
  * It returns true if the first value is less than or equal to the second and false if it is not.
  */
  if cookies > 10 {
    fmt.Println("That's too many cookies.")
  } else if cookies >= 5 {
    fmt.Println("That's a good amount of cookies.")
  } else if cookies <= 1 {
    fmt.Println("That's not very many cookies.")
  } else {
    fmt.Println("That's an okay amount of cookies.")
  }

  cookies = 15

  /*
  * && is an operator that takes two booleans and returns a boolean.
  * It returns true if both values are true and false if one or both values are false.
  *
  * || is an operator that takes two booleans and returns a boolean.
  * It returns true if one or both values are true and false if both values are false.
  */
  if cookies > 10 && cookies < 20 {
    fmt.Println("That's a lot of cookies.")
  } else if cookies > 10 || cookies < 20 {
    fmt.Println("That's an okay amount of cookies.")
  } else {
    fmt.Println("That's not very many cookies.")
  }

  cookies = 5

  /*
  * ! is an operator that takes a boolean and returns a boolean.
  * It returns true if the value is false and false if the value is true.
  */
  if !(cookies > 10) {
    fmt.Println("That's not a lot of cookies.")
  } else {
    fmt.Println("That's a lot of cookies.")
  }

  /*
  * % is an operator that takes two integers and returns an integer.
  * It returns the remainder of the first value divided by the second value.
  */
  if cookies % 2 == 0 {
    fmt.Println("That's an even number of cookies.")
  } else {
    fmt.Println("That's an odd number of cookies.")
  }
}