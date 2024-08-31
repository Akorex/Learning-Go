package main

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer

	for _, v := range s {
		r = f(r, v)
	}

	return r
}
