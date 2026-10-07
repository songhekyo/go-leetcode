package p1929concetanationofarray

import (
	"slices"
	"testing"
)

func TestConcetanationOfArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example_1", []int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConcatenation(tt.nums); !slices.Equal(got, tt.want) {
				t.Errorf("GetConcetanation(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
