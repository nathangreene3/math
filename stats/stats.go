package stats

import (
	"math"
	"sort"
)

// Frequency ...
type Frequency map[float64]int

// NewFreq ...
func NewFreq(vs ...float64) Frequency {
	freq := make(Frequency)
	for i := 0; i < len(vs); i++ {
		freq[vs[i]]++
	}

	return freq
}

// Count ...
func (f Frequency) Count() int {
	var c int
	for _, v := range f {
		c += v
	}

	return c
}

// Min ...
func (f Frequency) Min() float64 {
	if len(f) == 0 {
		panic("")
	}

	var (
		min   float64
		first = true
	)

	for k := range f {
		switch {
		case first:
			min = k
			first = false
		case k < min:
			min = k
		}
	}

	return min
}

// Max ...
func (f Frequency) Max() float64 {
	if len(f) == 0 {
		panic("")
	}

	var (
		max   float64
		first = true
	)

	for k := range f {
		switch {
		case first:
			max = k
			first = false
		case max < k:
			max = k
		}
	}

	return max
}

// Mean ...
func (f Frequency) Mean() float64 {
	var mean, n float64
	for k, v := range f {
		freq := float64(v)
		mean += k * freq
		n += freq
	}

	return mean / n
}

// Var ...
func (f Frequency) Var(mean float64) float64 {
	var sVar, n float64
	for k, v := range f {
		var dev, freq = k - mean, float64(v)
		sVar += dev * dev * freq
		n += freq
	}

	return sVar / (n - 1.0)
}

// StDev ...
func (f Frequency) StDev(mean float64) float64 {
	return math.Sqrt(f.Var(mean))
}

// Stats contains summary statistics data on a data set.
type Stats struct {
	count                            int
	min, max, mean, stdev, med, mode float64
	freq                             map[float64]int
}

// New summary statistics.
func New(vs ...float64) *Stats {
	s := Stats{
		count: len(vs),
		min:   vs[0],
		max:   vs[0],
		mean:  vs[0],
		mode:  vs[0],
		freq:  map[float64]int{vs[0]: 1},
	}

	distinctValues := make([]float64, 0, len(vs))
	for i := 1; i < len(vs); i++ {
		if _, ok := s.freq[vs[i]]; !ok {
			distinctValues = append(distinctValues, vs[i])
		}

		if vs[i] < s.min {
			s.min = vs[i]
		}

		if s.max < vs[i] {
			s.max = vs[i]
		}

		s.mean += vs[i]
		s.freq[vs[i]]++
		if s.mode < float64(s.freq[vs[i]]) {
			s.mode = vs[i]
		}
	}

	s.mean /= float64(len(vs))
	if 1 < len(vs) {
		for i := 0; i < len(vs); i++ {
			dev := s.mean - vs[i]
			s.stdev += dev * dev
		}

		s.stdev /= float64(len(vs) - 1)
	}

	sort.Float64s(distinctValues)
	var (
		j   int
		mid = len(vs) >> 1
	)

	for i := 0; i < len(distinctValues); i++ {
		// j is first index of value
		switch {
		case j < mid:
			if nextJ := j + s.freq[distinctValues[i]]; j < mid {
				j = nextJ
			} else {

			}
		case mid == j:
			if mid%2 == 0 {
				s.med = distinctValues[i]
			} else {
				s.med = (distinctValues[i-1] + distinctValues[i]) / 2.0
			}
		case mid < j:
			s.med = distinctValues[i]
		}
	}

	return &s
}

func new2(vs ...float64) *Stats {
	s := Stats{
		count: len(vs),
		min:   vs[0],
		max:   vs[0],
		mode:  vs[0],
		freq:  make(map[float64]int),
	}

	distinct := make([]float64, 0, len(vs))
	for i := 0; i < len(vs); i++ {
		if _, ok := s.freq[vs[i]]; !ok {
			distinct = append(distinct, vs[i])
		}

		s.freq[vs[i]]++
	}

	sort.Float64s(distinct)
	for i := 0; i < len(distinct); i++ {
		if distinct[i] < s.min {
			s.min = distinct[i]
		}

		if s.max < distinct[i] {
			s.max = distinct[i]
		}

		if s.mode != distinct[i] && s.freq[s.mode] < s.freq[distinct[i]] {
			s.mode = distinct[i]
		}

		s.mean += distinct[i] * float64(s.freq[distinct[i]])
	}

	s.mean /= float64(len(vs))
	for i := 0; i < len(distinct); i++ {
		diff := vs[i] - s.mean
		s.stdev += diff * diff
	}

	s.stdev /= float64(len(vs) - 1)
	s.stdev = math.Sqrt(s.stdev)

	if 1 < len(vs) {
		var (
			medIndex = len(vs) >> 1
			j        int
		)

		for i := 0; i < len(distinct); i++ {
			if medIndex == j {
				if len(vs)%2 == 0 {
					s.med = (distinct[i-1] + distinct[i]) / 2.0
				} else {
					s.med = distinct[i]
				}

				break
			}

			if medIndex < j {
				s.med = distinct[i]
				break
			}

			j += s.freq[distinct[i]]
		}
	} else {
		s.med = vs[0]
	}

	return &s
}

// Count returns the count.
func (s *Stats) Count() int {
	return s.count
}

// Freq returns the frequency a value occurs.
func (s *Stats) Freq(v float64) int {
	return s.freq[v]
}

// Max returns the maximum.
func (s *Stats) Max() float64 {
	return s.max
}

// Mean returns the mean.
func (s *Stats) Mean() float64 {
	return s.mean
}

// Med returns the median.
func (s *Stats) Med() float64 {
	return s.med
}

// Min returns the minimum.
func (s *Stats) Min() float64 {
	return s.min
}

// Mode returns the mode.
func (s *Stats) Mode() float64 {
	return s.mode
}

// Slice returns a sorted copy of the summarized data.
func (s Stats) Slice() []float64 {
	slc := make([]float64, 0, s.count)
	for k, v := range s.freq {
		for i := 0; i < v; i++ {
			slc = append(slc, k)
		}
	}

	sort.Float64s(slc)
	return slc
}

// StDev returns the standard deviation.
func (s *Stats) StDev() float64 {
	return s.stdev
}
