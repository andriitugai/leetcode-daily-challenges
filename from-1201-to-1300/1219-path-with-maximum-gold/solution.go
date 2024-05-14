type cell struct {
    row int
    col int
}

func getMaximumGold(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    seen := make(map[cell]bool)

    var dfs func(c cell) int 
    dfs  = func(c cell) int {
        row, col := c.row, c.col
        if row < 0 || row >= m || col < 0 || col >= n || seen[c] || grid[row][col] == 0 {
            return 0
        }
        seen[c] = true
        gl := dfs(cell{row: row, col: col - 1,})
        gr := dfs(cell{row: row, col: col + 1,})
        gd := dfs(cell{row: row + 1, col: col,})
        gu := dfs(cell{row: row - 1, col: col,})
        seen[c] = false

        return grid[row][col] + max(gl, max(gr, max(gd, gu)))
    }

    result := 0
    for row := 0; row < m; row ++ {
        for col := 0; col < n; col ++ {
            if grid[row][col] > 0 {
                c := cell{row: row, col: col,}
                if !seen[c] {
                    result = max(result, dfs(c))
                }
            }
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}