func numMagicSquaresInside(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    result := 0
    for r := 0; r < m - 2; r ++ {
        for c := 0; c < n - 2; c ++ {
            digits := make(map[int]bool)
            for dr := 0; dr < 3; dr ++ {
                for dc := 0; dc < 3; dc ++ {
                    if grid[r+dr][c+dc] < 10 && grid[r+dr][c+dc] > 0 {
                        digits[grid[r+dr][c+dc]] = true
                    }
                }
            }
            if len(digits) == 9 && 
                grid[r][c] + grid[r][c+1] + grid[r][c+2] == 15 && 
                grid[r+1][c] + grid[r+1][c+1] + grid[r+1][c+2] == 15 && 
                grid[r+2][c] + grid[r+2][c+1] + grid[r+2][c+2] == 15 && 
                grid[r][c] + grid[r+1][c] + grid[r+2][c] == 15 && 
                grid[r][c+1] + grid[r+1][c+1] + grid[r+2][c+1] == 15 && 
                grid[r][c+2] + grid[r+1][c+2] + grid[r+2][c+2] == 15 && 
                grid[r][c] + grid[r+1][c+1] + grid[r+2][c+2] == 15 && 
                grid[r][c+2] + grid[r+1][c+1] + grid[r+2][c] == 15 {
                    result += 1
                }
        }
    }
    return result
}