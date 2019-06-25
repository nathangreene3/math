package functional

// Mapper ...
type Mapper func(v interface{}) interface{}

// Filterer ...
type Filterer func(v interface{}) bool

// Reducer ...
type Reducer func(u, v interface{}) interface{}

// Map ...
func Map(S []interface{}, f Mapper) []interface{} {
	T := make([]interface{}, 0, len(S))
	for i := range S {
		T = append(T, f(S[i]))
	}

	return T
}

// Filter ...
func Filter(S []interface{}, f Filterer) []interface{} {
	T := make([]interface{}, 0, len(S))
	for _, v := range S {
		if f(v) {
			T = append(T, v)
		}
	}

	return T
}

// Reduce ...
func Reduce(S []interface{}, f Reducer) interface{} {
	n := len(S)
	if n == 0 {
		return nil
	}

	v := S[0]
	for i := 1; i < n; i++ {
		v = f(v, S[i])
	}

	return v
}