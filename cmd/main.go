package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(fizzbuzz(3, 5, 100))
}

func fizzbuzz(a, b, n int) string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = strconv.Itoa(i + 1)
	}

	for i := a - 1; i < n; i += a {
		s[i] = "Fizz"
	}

	for i := b - 1; i < n; i += b {
		s[i] = "Buzz"
	}

	ab := a * b
	for i := ab - 1; i < n; i += ab {
		s[i] = "FizzBuzz"
	}

	return strings.Join(s, "\n")
}
