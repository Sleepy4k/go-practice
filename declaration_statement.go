package main

import "fmt"

func main() {
  totalMoney := 25000

  /*
  * We can declare a variable inside if statement, and the variable will only be accessible inside the if statement.
  */
  if rob := totalMoney / 2; rob > 10000 {
    fmt.Println("Rob is rich.")
  } else {
    fmt.Println("Rob is not rich.")
  }

  /*
  * Not only if statement, we can also declare a variable inside switch statement, and the variable will only be accessible inside the switch statement.
  */
  switch rob := totalMoney / 2; {
    case rob > 10000:
      fmt.Println("Rob is rich.")
    case rob > 5000:
      fmt.Println("Rob is not rich.")
    default:
      fmt.Println("Rob is poor.")
  }
}
