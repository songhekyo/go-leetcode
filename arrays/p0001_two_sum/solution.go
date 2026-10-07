package p0001twosum

// 1. Two Sum
// https://leetcode.com/problems/two-sum/
//
// Diberikan array nums dan sebuah angka target. Kembalikan index dari dua
// angka yang jika dijumlahkan hasilnya sama dengan target.
//
// Setiap input dijamin punya tepat satu jawaban, dan elemen yang sama tidak
// boleh dipakai dua kali. Urutan index di jawaban bebas.
//
// Contoh:
//   nums = [2,7,11,15], target = 9 → [0,1]   (2 + 7 = 9)
//   nums = [3,2,4],     target = 6 → [1,2]   (2 + 4 = 6)
//   nums = [3,3],       target = 6 → [0,1]
//
// Batasan:
//   2 <= len(nums) <= 10^4
//   -10^9 <= nums[i] <= 10^9
//   -10^9 <= target <= 10^9
//
// Tantangan: bisakah kamu membuatnya lebih cepat dari O(n^2)?

func TwoSum(nums []int, target int) []int {
	mapped := make(map[int]int)

	for i, v := range nums {
		if j, ok := mapped[target-v]; ok {
			return []int{j, i}
		}
		mapped[v] = i
	}

	return []int{0, 0}
}
