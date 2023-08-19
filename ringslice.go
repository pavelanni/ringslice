package ringslice

type RingSliceInt struct {
	source *[]int
}

func NewInt(s *[]int) RingSliceInt {
	return RingSliceInt{source: s}
}

func (r RingSliceInt) Value(i int) int {
	s := *r.source
	newIndex := (i%len(s) + len(s)) % len(s)
	return s[newIndex]
}

func (r RingSliceInt) SetValue(i int, v int) {
	s := *r.source
	newIndex := (i%len(s) + len(s)) % len(s) // to process negative indices we have to add len(s) to the remainder and then again do %len(s)
	s[newIndex] = v
}
