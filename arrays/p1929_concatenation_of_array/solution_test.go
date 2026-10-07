package p1929concatenationofarray

import (
	"slices"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example_1", []int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
		{"example_2", []int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
		{"single_element", []int{7}, []int{7, 7}},
		{"all_same", []int{5, 5, 5}, []int{5, 5, 5, 5, 5, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetConcatenation(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Errorf("GetConcatenation(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// Solusi tidak boleh mengubah isi nums milik pemanggil.
func TestGetConcatenationDoesNotModifyInput(t *testing.T) {
	nums := make([]int, 3, 10) // sengaja diberi sisa kapasitas
	copy(nums, []int{1, 2, 3})

	GetConcatenation(nums)

	if got := nums[:cap(nums)][3:6]; !slices.Equal(got, []int{0, 0, 0}) {
		t.Errorf("array di belakang nums ikut berubah: %v", got)
	}
}
