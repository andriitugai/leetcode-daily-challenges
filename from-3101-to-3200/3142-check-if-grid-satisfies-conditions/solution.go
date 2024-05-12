func satisfiesConditions(grid [][]int) bool {
    m, n := len(grid), len(grid[0])

    for i := 0; i < m; i ++ {
        for j := 0; j < n; j ++ {
            if i < m - 1 && grid[i][j] != grid[i + 1][j] {
                return false
            }
            if j < n - 1 && grid[i][j] == grid[i][j + 1] {
                return false
            }
        }
    }
    return true
}