package flts

import (
	"testing"
)

func TestSet(t *testing.T) {
	t.Error(Map(New(1, 2, 3, 4, 5).Copy(), func(x float64) float64 { return 2.0 * x }))
	t.Error(Filter(New(1, 2, 3, 4, 5).Copy(), func(x float64) bool { return int64(x)%2 == 0 }))
	t.Error(Reduce(New(1, 2, 3, 4, 5).Copy(), func(x, y float64) float64 { return x + y }))
}
