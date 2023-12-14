func onesMinusZeros(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])

    rows := make([]int, m)
    cols := make([]int, n)

    for i := 0; i < m; i ++ {
        for j := 0; j < n; j ++ {
            if grid[i][j] == 1 {
                rows[i] += 1
                cols[j] += 1
            } else {
                rows[i] -= 1
                cols[j] -= 1
            }
        }
    }

    result := make([][]int, m)
    for i := 0; i < m; i ++ {
        result[i] = make([]int, n)
    }

    for i := 0; i < m; i ++ {
        for j := 0; j < n; j ++ {
            result[i][j] = rows[i] + cols[j]
        }
    }
    return result
}