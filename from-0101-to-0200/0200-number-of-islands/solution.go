func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    var dfs func(r, c int)
    dfs = func(r, c int) {
        if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == '0' {
            return
        }
        grid[r][c] = '0'
        dfs(r + 1, c)
        dfs(r - 1, c)
        dfs(r, c + 1)
        dfs(r, c - 1)
    }
    result := 0
    for i := 0; i < m; i ++ {
        for j := 0; j < n; j ++ {
            if grid[i][j] == '1' {
                result += 1
                dfs(i, j)
            }
        }
    }
    return result
}