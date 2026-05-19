func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)
	
	for _, i := range nums{
		count[i] ++
	}

	for num, cnt := range count {
		freq[cnt] = append(freq[cnt], num)
	} 

	var result []int

	for i:= len(freq) -1 ; i>=0; i--{
		result = append(result, freq[i]...)

		if len(result) == k {
			return result
		}

	}

	return result
}
