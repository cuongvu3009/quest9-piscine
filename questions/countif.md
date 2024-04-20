## Instructions

Write a function `CountIf` that returns the number of elements of a string slice for which the `f` function returns true.

### Expected function

```go
func CountIf(f func(string) bool, tab []string) int {

}

package main

import (
    "fmt"
    "piscine"
)
```

```go
func main() {
    tab1 := []string{"Hello", "how", "are", "you"}
    tab2 := []string{"This","1", "is", "4", "you"}

    answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
    answer2 := piscine.CountIf(piscine.IsNumeric, tab2)

    fmt.Println(answer1)
    fmt.Println(answer2)
}
```

```sh
$ go run .
0
2
$
