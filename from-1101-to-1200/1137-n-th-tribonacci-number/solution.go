func tribonacci(n int) int {
    tribo := make([]int, n + 3)
    tribo[0], tribo[1], tribo[2] = 0, 1, 1
    if n < 3 {
        return tribo[n]
    }

    for i := 3; i < n + 1; i ++ {
        tribo[i] = tribo[i-3] + tribo[i-2] + tribo[i-1]
    }
    return tribo[n]
}