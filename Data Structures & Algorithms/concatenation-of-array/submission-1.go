func getConcatenation(nums []int) []int {
    l := len(nums)

    if l ==0 {
        return []int{}
    }

    res := make([]int, 2*l)

    for i, _ := range nums {
        res[i], res[i+l] = nums[i], nums[i]
    }

    return res
}
