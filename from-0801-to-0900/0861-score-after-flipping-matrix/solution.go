func matrixScore(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    result := (1 << (n - 1)) * m

    for j := 1; j < n; j ++ {
        val := 1 << (n - 1 - j)
        setCount := 0
        for i := 0; i < m; i ++ {
            if grid[i][j] == grid[i][0] {
                setCount += 1
            }
        }
        result += max(setCount, m - setCount) * val
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}