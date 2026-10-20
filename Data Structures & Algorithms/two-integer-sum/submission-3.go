func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	prevMap := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if idx, ok := prevMap[diff]; ok {
			return []int{idx, i}
		}

		prevMap[n] = i
	}

	return []int{}

}
