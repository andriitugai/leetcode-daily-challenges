func hasIncreasingSubarrays(nums []int, k int) bool {
    maxGrowLen := 0
    curLen := 1
    prevLen := 0

    for i := 1; i < len(nums); i ++ {
        if nums[i] > nums[i - 1] {
            curLen += 1
        } else {
            prevLen = curLen
            curLen = 1
        }
        maxGrowLen = max(maxGrowLen, curLen / 2, min(curLen, prevLen))
    }
    return maxGrowLen >= k
}

func max(a, b, c int) int {
    if a > b {
        if a > c {
            return a
        }
        return c
    } else {
        if b > c {
            return b
        }
    }
    return c
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}