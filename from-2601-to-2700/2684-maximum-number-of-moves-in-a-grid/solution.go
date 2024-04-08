type Cell struct {
    row int
    col int
}

func maxMoves(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    memo := make(map[Cell]int)
    dirs := [][]int{[]int{0, 1}, []int{-1, 1}, []int{1, 1}}

    var dp func (r, c int) int
    dp = func (r, c int) int {
        cell := Cell{row: r, col: c}
        if v, ok := memo[cell]; ok {
            return v
        }
        maxPath := 0
        for _, dir := range dirs {
            nr, nc := r + dir[0], c + dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] > grid[r][c] {
                maxPath = max(maxPath, 1 + dp(nr, nc))
            }
        }
        memo[cell] = maxPath
        return maxPath
    }
    result := 0
    for r := 0; r < m; r ++ {
        result = max(result, dp(r, 0))
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}