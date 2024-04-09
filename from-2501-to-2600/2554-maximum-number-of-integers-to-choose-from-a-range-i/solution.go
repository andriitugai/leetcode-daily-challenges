func maxCount(banned []int, n int, maxSum int) int {
    isBanned := make(map[int]bool)
    for _, b := range banned {
        if b <= n {
            isBanned[b] = true
        }
    }
    sum, cnt := 0, 0
    for i := 1; i <= n; i ++ {
        if !isBanned[i] && i + sum <= maxSum {
            cnt += 1
            sum += i
        }
    }
    return cnt
}