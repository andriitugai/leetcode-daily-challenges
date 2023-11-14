func deleteGreatestValue(grid [][]int) int {
    m, n := len(grid), len(grid[0])

    for i := 0; i < m; i++ {
        sort.Ints(grid[i])
    }

    result := 0
    for j := 0; j < n; j ++ {
        maxVal := -1
        for i := 0; i < m; i ++ {
            if grid[i][j] > maxVal {
                maxVal = grid[i][j]
            }
        }
        result += maxVal
    }
    return result
}