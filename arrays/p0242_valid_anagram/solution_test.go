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
		{"example_2", "rat", "car", false},
		{"single_same_letter", "a", "a", true},
		{"identical", "listen", "listen", true},
		{"different_length", "a", "ab", false},
		{"t_shorter", "ab", "a", false},
		{"same_letters_different_count", "aacc", "ccac", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("ValidAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
