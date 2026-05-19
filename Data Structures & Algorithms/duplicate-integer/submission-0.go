func hasDuplicate(nums []int) bool {

    if len(nums) == 0 {
        return false
    }

    res := make(map[int]struct{})
    for _, n := range nums {
        if _, ok := res[n]; ok {
            return true
        }
        res[n] = struct{}{}
    }

    return false
    
}
