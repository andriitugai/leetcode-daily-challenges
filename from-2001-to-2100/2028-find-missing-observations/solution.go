func missingRolls(rolls []int, mean int, n int) []int {
    m := len(rolls)
    curSum := 0
    for _, roll := range rolls {
        curSum += roll
    }
    tgtSum := mean * (n + m)
    missedSum := tgtSum - curSum
    if missedSum < n || missedSum > n * 6 {
        return []int{}
    }

    result := make([]int, n)
    for i := 0; i < n; i ++ {
        result[i] = 1
    }
    missedSum -= n

    for i := 0; i < n; i ++ {
        result[i] += min(5, missedSum)
        missedSum -= 5
        if missedSum <= 0 {
            break
        }
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}