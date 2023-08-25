package main

import "fmt"

func main() {
  /*
  * We can declare a variable with var keyword, followed by the variable name and the type.
  * For example:
  * var name string
  *
  * We can also declare multiple variables in one line, for example:
  * var name, age string
  *
  * We can also declare a variable with initial value, for example:
  * var name string = "John"
  *
  * We can also declare multiple variables in one line with initial value, for example:
  * var name, age string = "John", "20"
  *
  * We can also declare a variable without var keyword, for example:
  * name := "John"
  *
  * We can also declare multiple variables in one line without var keyword, for example:
  * name, age := "John", "20"
  */
  var name string = "John"
  var age int = 20

  fmt.Println("My name is", name, "and I am", age, "years old.")

  name = "Doe"
  age = 30

  fmt.Println("My name is", name, "and I am", age, "years old.")

  name, age = "Jane", 25

  fmt.Println("My name is", name, "and I am", age, "years old.")

  var (
    firstName string = "John"
    lastName string = "Doe"
    myAge int = 20
  )

  fmt.Println("My name is", firstName, lastName, "and I am", myAge, "years old.")

  height := 170
  weight := 70

  fmt.Println("My height is", height, "cm and my weight is", weight, "kg.")

  height = 180
  weight = 80

  fmt.Println("My height is", height, "cm and my weight is", weight, "kg.")

  /*
  * We can also declare a constant with const keyword, followed by the constant name and the value.
  * For example:
  * const pi float32 = 3.14
  *
  * We can also declare multiple constants in one line, for example:
  * const pi float32 = 3.14, e float32 = 2.718
  *
  * We can also declare a constant without const keyword, for example:
  * pi := 3.14
  *
  * We can also declare multiple constants in one line without const keyword, for example:
  * pi, e := 3.14, 2.718
  */
  const pi float32 = 3.14
  const e float32 = 2.718

  fmt.Println("The value of pi is", pi, "and the value of e is", e)

  const (
    phi float32 = 1.618
    avogadro float32 = 6.022
  )

  fmt.Println("The value of phi is", phi, "and the value of avogadro is", avogadro)
}