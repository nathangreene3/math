package set2

import "testing"

// number is an test example of a comparable type.
type number int64

// Compare two numbers.
func (n number) Compare(c Comparable) int {
	switch {
	case n < c.(number):
		return -1
	case c.(number) < n:
		return 1
	default:
		return 0
	}
}

func TestSet(t *testing.T) {
	s := Generate(5, func(i int) Comparable { return number(i + 1) })
	t.Errorf("\nSet of [1, 5] = %v\n", s)
	t.Errorf("\nSum of [1, 5] = %d\n", s.Reduce(func(x, y Comparable) Comparable { return x.(number) + y.(number) }))
	t.Errorf("\nSum of [1, 5] = %d\n", s.Fold(number(0), func(x interface{}, v Comparable) interface{} { return x.(number) + v.(number) }))
	t.Errorf("\nMin of [1, 5] = %d\n", s.Min())
	t.Errorf("\nMax of [1, 5] = %d\n", s.Max())
	t.Errorf("\nSubset [2, 4] of [1, 5] = %v\n", s.SubsetRangeEqual(number(2), number(4)).Slice())
	t.Errorf("\nSubset {1, 5} of [1, 5] = %v\n", s.SubsetOmitRange(number(2), number(4)).Slice())
	t.Errorf("\nSubset of evens in [1, 5] = %v\n", s.Filter(func(v Comparable) bool { return v.(number)%2 == 0 }).Slice())
	t.Errorf("\nUnion of evens and odds = %v\n", s.Filter(func(v Comparable) bool { return v.(number)%2 == 0 }).Union(s.Filter(func(v Comparable) bool { return v.(number)%2 != 0 })).Slice())
	t.Errorf("\nSubset of odds in [1, 5] = %v\n", s.Remove(number(2), number(4)).Slice())
	t.Errorf("\nEvens are a subset of [1, 5] = %t\n", s.ContainsSubset(s.Filter(func(v Comparable) bool { return v.(number)%2 == 0 })))
}
