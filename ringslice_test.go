package ringslice

import (
	"testing"
)

func TestRingSlice(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	ringInt := New(&ints)

	tests := []struct {
		index    int
		expected int
	}{
		{0, 1},
		{1, 2},
		{5, 1},
		{-1, 5},
		{-6, 5},
	}

	for _, tt := range tests {
		if got := ringInt.Value(tt.index); got != tt.expected {
			t.Errorf("Value(%d) = %d; want %d", tt.index, got, tt.expected)
		}
	}

	ringInt.SetValue(5, 10)
	if ringInt.Value(0) != 10 {
		t.Errorf("Value after SetValue not updated correctly")
	}
}

func TestRingSliceWithString(t *testing.T) {
	strs := []string{"a", "b", "c"}
	ringStr := New(&strs)

	if got := ringStr.Value(3); got != "a" {
		t.Errorf("Value(3) = %s; want 'a'", got)
	}

	ringStr.SetValue(-4, "z")
	if ringStr.Value(2) != "z" {
		t.Errorf("Value after SetValue not updated correctly")
	}
}

func TestRingSliceWithEmptySlice(t *testing.T) {
	empty := []int{}
	ringEmpty := New(&empty)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic with empty slice but didn't get one")
		}
	}()

	// This should panic due to modulo with zero
	ringEmpty.Value(0)
}
