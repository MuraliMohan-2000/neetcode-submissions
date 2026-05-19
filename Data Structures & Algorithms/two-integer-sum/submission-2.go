func twoSum(nums []int, target int) []int {

    if len(nums) == 0 {
        return []int{}
    }

    diffMap := make(map[int]int)

    for i, n := range nums {

        diff := target-n

        if idx, ok := diffMap[diff]; ok {
            return []int{idx, i}
        }

        diffMap[n] = i
    }

return []int{}   
}
