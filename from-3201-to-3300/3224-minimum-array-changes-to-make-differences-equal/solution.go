func minChanges(nums []int, k int) int {
    diffsCount := make([]int, k + 1)
    maxAttain := make([]int, k + 1)
    n := len(nums)
    var a, b int

    for i := 0; i < n / 2; i ++ {
        a, b = nums[i], nums[n - i - 1]
        diffsCount[abs(a - b)] += 1
        maxDiff := max(max(a, k - a), max(b, k - b))
        maxAttain[maxDiff] += 1
    }

    result := n
    curSum := 0
    for i := k; i >= 0; i -- {
        curSum += maxAttain[i]
        countGroupRight := curSum - diffsCount[i]
        countGroupLeft := n / 2 - countGroupRight - diffsCount[i]

        ops := countGroupLeft * 2 + countGroupRight
        result = min(result, ops)
    }
    return result
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}