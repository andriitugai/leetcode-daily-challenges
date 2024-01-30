func countOfPairs(n int, x int, y int) []int {
    if x > y {
        x, y = y, x
    }
    result := make([]int, n)
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }
    abs := func(a int) int {
        if a < 0 {
            return -a
        }
        return a
    }

    for i := 1; i <= n; i ++ {
        for j := i + 1; j <= n; j ++ {
            v := min(j - i, 1 + abs(i - x) + abs(j - y))
            result[v - 1] += 2
        }
    }

    return result
}