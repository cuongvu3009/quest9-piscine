## Instructions

Write a program called `doop`.

The program has to be used with three arguments:

1. A value
2. An operator, one of: +, -, /, *, %
3. Another value

In case of an invalid operator, value, number of arguments, or an overflow, the program prints nothing.

The program has to handle the modulo and division operations by 0 as shown in the output examples below.

## Usage

```sh
$ go run .
$ go run . 1 + 1 | cat -e
2
$
$ go run . hello + 1
$ go run . 1 p 1
$ go run . 1 / 0 | cat -e
No division by 0
$
$ go run . 1 % 0 | cat -e
No modulo by 0
$
$ go run . 9223372036854775807 + 1
$ go run . -9223372036854775809 - 3
$ go run . 9223372036854775807 "*" 3
$ go run . 1 "*" 1
1
$ go run . 1 "*" -1
-1
$
