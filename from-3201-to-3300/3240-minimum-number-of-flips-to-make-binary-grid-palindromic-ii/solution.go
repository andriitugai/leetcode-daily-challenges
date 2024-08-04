func minFlips(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    result := 0
    single, double := 0, 0

    for i := 0; i < m / 2; i ++ {
        for j := 0; j < n / 2; j ++ {
            ones := grid[i][j] + grid[i][n-1-j] + grid[m-1-i][j] + grid[m-1-i][n-1-j]
            result += min(ones, 4 - ones)
        }
        if n % 2 == 1 {
            ones := grid[i][n / 2] + grid[m - 1 - i][n / 2]
            if ones == 1 {
                single += 1
            } else if ones == 2 {
                double += 1
            }
        }
    }
    if m % 2 == 1 {
        for j := 0; j < n / 2; j ++ {
            ones := grid[m / 2][j] + grid[m / 2][n - 1 - j]
            if ones == 1 {
                single += 1
            } else if ones == 2 {
                double += 1
            }
        }
        if n % 2 == 1 {
            result += grid[m / 2][n / 2]
        }
    }
    if double % 2 == 0 || single > 0 {
        return result + single
    }
    return result + 2
}