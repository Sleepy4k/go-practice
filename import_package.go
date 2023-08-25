package main

/*
* To import a package, we use the import keyword, followed by the name of the package.
* In this case, we are importing the fmt package. for example:
* import "fmt"
*
* But in case we want to import multiple packages, we can do it like this:
* import (
*   "fmt"
*   "math"
* )
*
* And we can put an alias for the package name, for example:
* import (
*   "fmt"
*   m "math"
* )
*/
import (
  "fmt"
  t "time"
)

func main() {
  /*
  * After add the package alias name, we can use it to call the function.
  * For example, we can use t.Now() to call the Now() function from time package.
  */
  fmt.Println(t.Now())
}