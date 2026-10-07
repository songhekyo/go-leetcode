package p0303rangesumqueryimmutable

import "testing"

func TestSumRange(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)

	tests := []struct {
		name        string
		left, right int
		want        int
	}{
		{"example_1", 0, 2, 1},
		{"example_2", 2, 5, -1},
		{"example_3", 0, 5, -3},
		{"single_element", 3, 3, -5},
		{"first_element", 0, 0, -2},
		{"last_element", 5, 5, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := obj.SumRange(tt.left, tt.right); got != tt.want {
				t.Errorf("SumRange(%d, %d) = %d, want %d", tt.left, tt.right, got, tt.want)
			}
		})
	}
}
