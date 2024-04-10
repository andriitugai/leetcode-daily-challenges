func maxJump(stones []int) int {
    n := len(stones)
    if n == 2 {
        return stones[1] - stones[0]
    }
    if n % 2 == 0 {
        stones = append(stones, stones[n - 1] + 1)
        n += 1
    }
    result := 0
    for i := 0; i < n - 2; i += 2 {
        result = max(result, stones[i + 2] - stones[i])
    }
    for i := 1; i < n - 2; i += 2 {
        result = max(result, stones[i + 2] - stones[i])
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}