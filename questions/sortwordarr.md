## Instructions

Write a function `SortWordArr` that sorts by ASCII (in ascending order) a string slice.

### Expected function

```go
func SortWordArr(a []string) {

}

package main

import (
    "fmt"
    "piscine"
)
```
```go
func main() {
    result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
    piscine.SortWordArr(result)

    fmt.Println(result)
}
```
```sh
$ go run .
[1 2 3 A B C a b c]
$
