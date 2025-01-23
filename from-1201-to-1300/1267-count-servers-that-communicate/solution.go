func countServers(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    rows, cols := make([]int, m), make([]int, n)
    servers := 0
    for i := 0; i < m; i ++ {
        for j := 0; j < n; j ++ {
            if grid[i][j] == 1 {
                servers += 1
                rows[i] += 1
                cols[j] += 1
            }
        }
    }

    for i := 0; i < m; i ++ {
        for j := 0; j < n; j ++ {
            if grid[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
                servers -= 1
            }
        }
    }

    return servers
}