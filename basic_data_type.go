package main

import "fmt"

func main() {
  /*
  * There are several basic data types in Go:
  * - bool
  * - string
  * - int
  * - int8
  * - int16
  * - int32
  * - int64
  * - uint
  * - uint8
  * - uint16
  * - uint32
  * - uint64
  * - uintptr
  * - byte
  * - rune
  * - float32
  * - float64
  * - complex64
  * - complex128
  */
  var isTrue bool = true
  var name string = "John"
  var age int = 20
  var max uint = 1000
  var height float32 = 170.5

  fmt.Println("isTrue is", isTrue)
  fmt.Println("name is", name)
  fmt.Println("age is", age)
  fmt.Println("max is", max)
  fmt.Println("height is", height)

  /*
  * We can update the value of a variable with some operators, for example:
  * - +=
  * - -=
  * - *=
  * - /=
  * - %=
  */
  weight := 70

  age += 10
  max -= 100
  height *= 2
  weight /= 2

  fmt.Println("age is", age)
  fmt.Println("max is", max)
  fmt.Println("height is", height)
  fmt.Println("weight is", weight)
}