package p1929concetanationofarray

func GetConcatenation(nums []int) []int {
	arr := make([]int, 0, len(nums)*2)
	arr = append(arr, nums...)
	arr = append(arr, nums...)

	return arr
}
