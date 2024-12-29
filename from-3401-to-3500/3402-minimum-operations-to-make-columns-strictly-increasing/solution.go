func minimumOperations(grid [][]int) int {
    result := 0
    for j := 0; j < len(grid[0]); j ++ {
        prev := grid[0][j]
        for i := 1; i < len(grid); i ++ {
            if grid[i][j] <= prev {
                result += prev - grid[i][j] + 1
                grid[i][j] = prev + 1
            }
            prev = grid[i][j]
        }
    }
    return result
}