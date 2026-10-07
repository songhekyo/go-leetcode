package p0242validanagram

func ValidAnagram(s string, t string) bool {
	mapped_of_s := make(map[rune]int)

	for _, v := range s {
		mapped_of_s[v]++
	}

	for _, v := range t {
		mapped_of_s[v]--
		if mapped_of_s[v] < 0 {
			return false
		}
	}

	return true
}
