func countCharacters(words []string, chars string) int {
	getCount := func(s string) map[rune]int {
		cnt := make(map[rune]int)
		for _, c := range s {
			cnt[c] += 1
		}
		return cnt
	}
	result := 0
	lc := len(chars)
	charsCnt := getCount(chars)
	for _, word := range words {
		if len(word) <= lc {
			isOk := true
			curCnt := getCount(word)
			for k, v := range curCnt {
				if charsCnt[k] < v {
					isOk = false
					break
				}
			}
			if isOk {
				result += len(word)
			}
		}
	}
	return result
}