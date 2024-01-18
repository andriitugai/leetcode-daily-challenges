func climbStairs(n int) int {
    s1, s2 := 1, 1

    for i := 1; i < n; i ++ {
        s1, s2 = s1 + s2, s1
    }
    return s1
}