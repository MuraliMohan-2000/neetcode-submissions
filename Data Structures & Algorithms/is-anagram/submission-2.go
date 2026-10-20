func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	count := [26]int{}
	for i, _ := range s{
		count[s[i] - 'a']++
		count[t[i] - 'a']--
	}

	for i, _ := range count{
		if count[i] != 0 {
			return false
		}
	}

	return true
}
