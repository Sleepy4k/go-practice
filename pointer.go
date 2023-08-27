package main

import "fmt"

/*
* If we want to change the value of a variable inside a function, we can pass the memory address of the variable to the function.
*/
func changeValue(word *string) {
  *word = "Breaddit"
}

func main() {
  /*
  * Pointer is a variable that stores the memory address of another variable.
  */
  name := "John"

  fmt.Println("The memory address of name variable is", &name)

  /*
  * We can also declare a pointer variable by using * operator.
  */
  var namePointer *string = &name

  fmt.Println("The value of namePointer variable is", namePointer)

  /*
  * We can get the value of a variable by using * operator.
  */
  fmt.Println("The value of namePointer variable is", *namePointer)

  /*
  * We can also change the value of a variable by using * operator.
  */
  *namePointer = "Doe"

  fmt.Println("The value of name variable is", name)

  /*
  * We can also declare a pointer variable by using new keyword.
  */
  var agePointer *int = new(int)

  fmt.Println("The value of agePointer variable is", agePointer)

  /*
  * We can also change the value of a variable by using * operator.
  */
  *agePointer = 20

  fmt.Println("The value of agePointer variable is", *agePointer)

  /*
  * Change the value of a variable inside a function.
  */
  changeValue(&name)

  fmt.Println("The value of name variable is", name)
}