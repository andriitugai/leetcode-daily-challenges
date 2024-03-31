func countSubarrays(nums []int, minK int, maxK int) int64 {
    bad, mi, ma := -1, -1, -1
    var result int64 = 0

    for i, n := range nums {
        if minK > n || n > maxK {
            bad = i
        }

        if minK == n {
            mi = i
        }
        if maxK == n {
            ma = i
        }
        start := min(mi, ma)
        if start > bad {
            result += int64(start - bad)
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