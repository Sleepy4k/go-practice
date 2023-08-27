package main

import "fmt"

func main() {
  step1 := "Breathe in..."
  step2 := "Breathe out..."

  /*
  * Sprintln is a function that takes any number of arguments of any type and returns a string.
  * It takes the values and converts them to strings, then adds spaces between the values and a newline character at the end.
  */
  meditation := fmt.Sprintln(step1, step2)

  wish := "I wish i had a %v and %v."
  gift := "cat"
  secondGift := "dog"

  /*
  * Sprintf is a function that takes a string as the first argument, then any number of values.
  * It returns a string that is formatted as the original string with the values substituted in.
  * value types are converted to strings and placed in the string where the %v placeholders are.
  */
  var result string

  result = fmt.Sprintf(wish, gift, secondGift)

  fmt.Println(meditation)
  fmt.Println(result)
}