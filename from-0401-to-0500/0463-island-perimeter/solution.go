type Cell struct {
    row int
    col int
}

func islandPerimeter(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    seen := make(map[Cell]bool)

    var dfs func(row, col int) int
    dfs = func(row, col int) int {
        if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == 0 {
            return 1
        }
        cell := Cell{row: row, col: col}
        if seen[cell] {
            return 0
        }
        seen[cell] = true
        perimeter := dfs(row + 1, col)
        perimeter += dfs(row - 1, col)
        perimeter += dfs(row, col + 1)
        perimeter += dfs(row, col - 1)
        return perimeter
    }
    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            if grid[r][c] == 1 {
                return dfs(r, c)
            }
        }
    }
    return 0
}