package main

import "fmt"

type T int

func (t *T) add(u T) *T {
	*t += u
	return t
}

func (t *T) sub(u T) *T {
	*t -= u
	return t
}

func main() {
	var t T
	t.add(1).sub(2)
	fmt.Println(t)

	t = 0
	t = *t.add(1).sub(2)
	fmt.Println(t)
}
