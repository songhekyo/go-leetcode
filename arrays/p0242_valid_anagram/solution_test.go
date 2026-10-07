package p0242validanagram

import "testing"

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"example_1", "anagram", "nagaram", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("ContainsDuplicate(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)

			}
		})
	}
}
