func getConcatenation(nums []int) []int {
    n := len(nums)
    var res []int
    if n < 0 {
        return res
    }
    res = make([]int, 2*n)
    for i, _ := range nums{
    res[i], res[i+n] = nums[i], nums[i]
    }

    return res
    
}
