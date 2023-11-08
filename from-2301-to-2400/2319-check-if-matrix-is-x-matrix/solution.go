func checkXMatrix(grid [][]int) bool {
    n := len(grid)
    for r := 0; r < n; r ++ {
        for c := 0; c < n; c ++ {
            if r == c || r == n - c - 1{
                if grid[r][c] == 0 {
                    return false
                }
            } else {
                if grid[r][c] != 0 {
                    return false
                }
            }
        }
    }
    return true
}