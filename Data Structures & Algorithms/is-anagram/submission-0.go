func isAnagram(s string, t string) bool {
    
    if len(s) != len(t){
        return false
    }

    sMap, tMap := make(map[rune]int, len(s)), make(map[rune]int, len(t))

    for _, c := range s {
        sMap[c] += 1
    }

    for _, ch := range t {
        tMap[ch] += 1
    }

    for k, v := range sMap{
        count, ok  := tMap[k]
        if !ok {
            return false
        } else if (count != v){
            return false
        }
    }

    return true
}
