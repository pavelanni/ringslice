package ringslice

type RingSlice[T any] struct {
	source *[]T
}

func New[T any](s *[]T) RingSlice[T] {
	return RingSlice[T]{source: s}
}

func (r RingSlice[T]) Value(i int) T {
	s := *r.source
	newIndex := (i%len(s) + len(s)) % len(s)
	return s[newIndex]
}

func (r RingSlice[T]) SetValue(i int, v T) {
	s := *r.source
	newIndex := (i%len(s) + len(s)) % len(s) // to process negative indices we have to add len(s) to the remainder and then again do %len(s)
	s[newIndex] = v
}
