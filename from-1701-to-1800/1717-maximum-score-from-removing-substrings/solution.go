func maximumGain(s string, x int, y int) int {
	removePairs := func(pair string, score int) int {
		result := 0
		stack := []byte{}
		for i := 0; i < len(s); i ++ {
            if s[i] == pair[1] && len(stack) > 0 && stack[len(stack) - 1] == pair[0] {
                stack = stack[:len(stack) - 1]
                result += score
            } else {
                stack = append(stack, s[i])
            }
		}
        s = string(stack)
        return result
	}

    result := 0
    p1, p2 := "ab", "ba"
    s1, s2 := x, y
    if s2 > s1 {
        p1, p2 = p2, p1
        s1, s2 = s2, s1
    }
    result += removePairs(p1, s1)
    result += removePairs(p2, s2)
    return result
}