func groupAnagrams(strs []string) [][]string {

    if len(strs) == 0 {
        return [][]string{{""}}
    }
    
    frequencyMap := make(map[[26]int][]string)

    for _, str := range strs {
        freq := [26]int{}
        for _, c := range str {
            freq[c-'a'] += 1
        }

        frequencyMap[freq] = append(frequencyMap[freq], str)
    }

    var res [][]string

    for _, v := range frequencyMap {
        res = append(res, v)
    }

    return res

     

}
