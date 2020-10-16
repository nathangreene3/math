package bitmask0

import (
	"fmt"
	"testing"
)

func TestBitmask(t *testing.T) {
	a := New().Set(10)
	fmt.Printf("%s\n", a)

	t.Fatal()
}
