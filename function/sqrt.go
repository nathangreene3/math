package function

func sqrtNewton(x float64, n int) float64 {
	var r0, r1 float64 = 1, 0
	for ; 0 < n; n-- {
		r1 = (r0 + x/r0) / 2
		r0 = r1
	}

	return r1
}

func sqrtNewton2(x float64) float64 {
	var r0, r1 float64 = 1, 0
	for d := r1 - r0; d != 0; {
		r1 = (r0 + x/r0) / 2
		d = r1 - r0
		r0 = r1
	}

	return r1
}
