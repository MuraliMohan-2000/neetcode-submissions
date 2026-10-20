func getConcatenation(nums []int) []int {
	l := len(nums)

	if l <= 0 {
		return []int{}
	}

	ans := make([]int, 2*l)

	for i, _ := range nums{
		ans[i], ans[i+l] = nums[i], nums[i]
	}

	return ans
}
