package main

import "fmt"

/*
* Function is a block of code that can be executed multiple times.
* To declare a function, we use the func keyword, followed by the function name, and the function body.
* For example:
* func functionName() {
*   fmt.Println("Hello World!")
* }
*/
func makeSimpleFunction() {
  fmt.Println("Simple function.")
}

/*
* We can also declare a function with parameters.
* For example:
* func functionName(parameter1 string, parameter2 int) {
*   fmt.Println(parameter1, parameter2)
* }
*/
func makeFunctionWithParameters(name string, age int) {
  fmt.Println("My name is", name, "and I am", age, "years old.")
}

/*
* We can also declare a function with return value.
* For example:
* func functionName() string {
*   return "Hello World!"
* }
*/
func makeFunctionWithReturnValue() string {
  return "Hello World!"
}

/*
* We can also declare a function with multiple return values.
* For example:
* func functionName() (string, int) {
*   return "John", 20
* }
*/
func makeFunctionWithMultipleReturnValues() (string, int) {
  return "John", 20
}

/*
* We can simplify parameter data type declaration by using the same data type for all parameters.
* For example:
* func functionName(parameter1, parameter2, parameter3 string) {
*   fmt.Println(parameter1, parameter2, parameter3)
* }
*/
func makeFunctionWithSameParameterDataType(name, address, phoneNumber string) {
  fmt.Println("My name is", name, "and I live in", address, "and my phone number is", phoneNumber)
}

/*
* Inside function body, we can use async method using defer keyword.
* For example:
* func functionName() {
*   defer fmt.Println("World!")
*   fmt.Println("Hello")
* }
*
* The output will be:
* Hello
* World!
*/
func makeFunctionWithAsyncMethod() {
  defer fmt.Println("World!")
  fmt.Println("Hello")
}

func main() {
  makeSimpleFunction()

  makeFunctionWithParameters("John", 20)

  fmt.Println(makeFunctionWithReturnValue())

  name, age := makeFunctionWithMultipleReturnValues()

  fmt.Println("My name is", name, "and I am", age, "years old.")

  makeFunctionWithSameParameterDataType("John", "New York", "123456789")

  makeFunctionWithAsyncMethod()
}