package p0242validanagram

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/
//
// Diberikan dua string s dan t. Kembalikan true jika t adalah anagram dari s,
// dan false jika bukan.
//
// Anagram adalah kata yang dibentuk dengan menyusun ulang semua huruf dari
// kata lain, dengan setiap huruf dipakai tepat sebanyak jumlahnya di kata asal.
//
// Contoh:
//   s = "anagram", t = "nagaram" → true
//   s = "rat",     t = "car"     → false
//
// Batasan:
//   1 <= len(s), len(t) <= 5 * 10^4
//   s dan t hanya berisi huruf kecil a-z
//
// Tantangan: bagaimana jika input bisa berisi karakter Unicode (misalnya
// emoji atau huruf beraksen)? Apakah solusimu masih benar?

func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	anag := make(map[rune]int)

	for _, v := range s {
		anag[v]++
	}

	for _, v := range t {
		anag[v]--
		if anag[v] < 0 {
			return false
		}
	}

	return true
}
