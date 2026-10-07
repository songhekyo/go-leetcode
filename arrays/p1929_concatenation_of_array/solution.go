package p1929concatenationofarray

// 1929. Concatenation of Array
// https://leetcode.com/problems/concatenation-of-array/
//
// Diberikan array nums dengan panjang n. Buat array ans dengan panjang 2n
// di mana ans[i] == nums[i] dan ans[i + n] == nums[i] untuk 0 <= i < n.
//
// Contoh:
//   nums = [1,2,1]   → [1,2,1,1,2,1]
//   nums = [1,3,2,1] → [1,3,2,1,1,3,2,1]
//
// Batasan:
//   1 <= n <= 1000
//   1 <= nums[i] <= 1000

func GetConcatenation(nums []int) []int {
	result := make([]int, 0, len(nums)*2)

	result = append(result, nums...)
	result = append(result, nums...)

	return result
}
