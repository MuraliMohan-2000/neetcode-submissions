func isAnagram(s string, t string) bool {

	if len(s) != len(t){
		return false
	}

	freqS := make(map[rune]int)
	freqT := make(map[rune]int)


	for i, c := range s {
		freqS[c] += 1
		freqT[rune(t[i])] += 1
	}

	for k1, v1 := range freqS {
		if v2 := freqT[k1]; v1 != v2 {
			return false
		}
	}

	return true

}
