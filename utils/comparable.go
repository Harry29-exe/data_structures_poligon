package comparator

// Compare returns < 0 if v1 < v2, 0 if v1 == v2, > 0 if v1 > v2
type Compare[T any] = func(T, T) int8

type Primitive interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~string
}

func ComparePrimitive[P Primitive](v1, v2 P) int8 {
	if v1 < v2 {
		return -1
	} else if v1 == v2 {
		return 0
	}

	return 1
}
