func hasDuplicate(nums []int) bool {

	if len(nums) == 0 {
		return false
	}

	prevMap := make(map[int]bool)

	for i, _ := range nums {
		if _, ok := prevMap[nums[i]]; ok {
			return true
		}

		prevMap[nums[i]] = true
	}

	return false

}
