func hasDuplicate(nums []int) bool {
    if len(nums) == 0 {
        return false
    }

    prevMap := make(map[int]bool)

    for _, i := range nums {
        if _, ok := prevMap[i]; ok{
            return true
        }

        prevMap[i] = true
    }

    return false
}
