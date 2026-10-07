package p0001twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"example_1", []int{1, 3, 4}, 5, []int{0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.nums, tt.target); !slices.Equal(got, tt.want) {
				t.Errorf("ContainsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)

			}
		})
	}
}
