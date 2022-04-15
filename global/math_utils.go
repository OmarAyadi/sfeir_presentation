package global

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
