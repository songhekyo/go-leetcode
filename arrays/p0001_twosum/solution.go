package p0001twosum

func TwoSum(nums []int, target int) []int {
	mapped := make(map[int]int)

	for i, v := range nums {
		if j, ok := mapped[target-v]; ok {
			return []int{j, i}
		}
		mapped[v] = i
	}

	return nil
}
