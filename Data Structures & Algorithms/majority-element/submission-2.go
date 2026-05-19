func majorityElement(nums []int) int {

    if len(nums) <= 1{
        return nums[0]
    }

    halfLength := len(nums)/2

    freqMap := make(map[int]int)
    for _, n := range nums {
        if count, ok := freqMap[n]; ok {
            if count >= halfLength {
                return n
            }
        }

        freqMap[n] += 1
    }

    return 0
    
}
