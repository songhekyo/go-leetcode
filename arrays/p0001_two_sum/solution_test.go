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
		{"example_1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example_2", []int{3, 2, 4}, 6, []int{1, 2}}, // jangan pakai index 0 dua kali (3 + 3)
		{"example_3", []int{3, 3}, 6, []int{0, 1}},
		{"negative_numbers", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{"zeros", []int{0, 4, 3, 0}, 0, []int{0, 3}},
		{"answer_at_end", []int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)

			// urutan jawaban bebas, jadi urutkan dulu sebelum dibandingkan
			sorted := slices.Clone(got)
			slices.Sort(sorted)

			if !slices.Equal(sorted, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
