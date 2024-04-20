## Instructions

Write a function `ForEach` that, for an int slice, applies a function on each element of that slice.

### Expected function

```go
func ForEach(f func(int), a []int) {

}
```

```go
package main

import "piscine"

func main() {
    a := []int{1, 2, 3, 4, 5, 6}
    piscine.ForEach(piscine.PrintNbr, a)
}
```

```sh
$ go run .
123456
$
```