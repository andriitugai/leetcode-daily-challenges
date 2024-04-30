func numberOfRightTriangles(grid [][]int) int64 {
    m, n := len(grid), len(grid[0])
    rCounts, cCounts := make([]int, m), make([]int, n)
    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            if grid[r][c] == 1 {
                rCounts[r] += 1
                cCounts[c] += 1
            }
        }
    }
    var result int64 = 0
    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            if grid[r][c] == 1 {
                result += int64(rCounts[r] - 1) * int64(cCounts[c] - 1)
            }
        }
    }
    return result
}