package stats

// Stats contains summary statistics data on a data set.
type Stats struct {
	count                            int
	min, max, mean, stdev, med, mode float64
	freq                             map[float64]int
}

// New ...
func New(vs ...float64) Stats {
	s := Stats{
		count: len(vs),
		freq:  make(map[float64]int),
	}

	if s.count == 0 {
		return s
	}

	s.med = vs[len(vs)>>1]
	if 1 < len(vs) && len(vs)%2 == 0 {
		s.med = (s.med + vs[len(vs)>>1+1]) / 2
	}

	s.min = vs[0]
	s.max = vs[0]
	s.mean = vs[0]
	s.mode = vs[0]
	for i := 1; i < len(vs); i++ {
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

	return s
}

// Count returns the count.
func (s Stats) Count() int {
	return s.count
}

// Freq returns the frequency a value occurs.
func (s Stats) Freq(v float64) int {
	return s.freq[v]
}

// Max returns the maximum.
func (s Stats) Max() float64 {
	return s.max
}

// Mean returns the mean.
func (s Stats) Mean() float64 {
	return s.mean
}

// Med returns the median.
func (s Stats) Med() float64 {
	return s.med
}

// Min returns the minimum.
func (s Stats) Min() float64 {
	return s.min
}

// Mode returns the mode.
func (s Stats) Mode() float64 {
	return s.mode
}

// StDev returns the standard deviation.
func (s Stats) StDev() float64 {
	return s.stdev
}
