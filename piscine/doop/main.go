package main

import (
	"os"
)

func atof(s string) (float64, bool) {
	var v float64
	var sign float64 = 1
	var dec float64 = 1
	var after decimalState
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			if after == decimal {
				dec *= 0.1
				v += float64(c-'0') * dec
			} else {
				v = v*10 + float64(c-'0')
			}
		case c == '.':
			if after == noDecimal {
				after = decimal
			} else {
				return v, false
			}
		case c == '-':
			if i == 0 {
				sign = -1
			} else {
				return v, false
			}
		default:
			return v, false
		}
	}
	return v * sign, true
}

type decimalState int

const (
	noDecimal decimalState = iota
	decimal
)

func doOperation(value1 float64, operator string, value2 float64) (float64, bool) {
	switch operator {
	case "+":
		return value1 + value2, true
	case "-":
		return value1 - value2, true
	case "*":
		return value1 * value2, true
	case "/":
		if value2 == 0 {
			return 0, false
		}
		return round(value1 / value2), true
	case "%":
		if value2 == 0 {
			return 0, false
		}
		return float64(int64(value1) % int64(value2)), true
	default:
		return 0, false
	}
}
func round(x float64) float64 {
	return float64(int64(x))
}
func main() {
	if len(os.Args) != 4 {
		return
	}
	value1, valid1 := atof(os.Args[1])
	operator := os.Args[2]
	value2, valid2 := atof(os.Args[3])
	if value1 > 9223372036854775807 || value2 > 9223372036854775807 {
		return
	}
	if !valid1 || !valid2 || !isValidOperator(operator) {
		return
	}
	result, valid := doOperation(value1, operator, value2)
	if valid {
		printInt(result)
	} else {
		switch operator {
		case "/":
			if value1 == 9223372036854775807 && value2 == 0 {
				return
			}
			os.Stdout.Write([]byte("No division by 0\n"))
		case "%":
			if value2 == 0 {
				os.Stdout.Write([]byte("No modulo by 0\n"))
			}
		}
	}
}
func printInt(i float64) {
	var buf [50]byte
	var n int
	// Handle negative numbers
	if i < 0 {
		os.Stdout.Write([]byte{'-'})
		i = -i
	}
	// Extract digits
	for i >= 10 {
		rem := int(i) % 10
		buf[n] = byte(rem) + '0'
		n++
		i /= 10
	}
	buf[n] = byte(i) + '0'
	n++
	// Print in reverse order
	for j := n - 1; j >= 0; j-- {
		os.Stdout.Write([]byte{buf[j]})
	}
	os.Stdout.Write([]byte{'\n'})
}
func isValidOperator(operator string) bool {
	switch operator {
	case "+", "-", "*", "/", "%":
		return true
	default:
		return false
	}
}
