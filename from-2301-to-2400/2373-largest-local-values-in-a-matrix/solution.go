func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    getMax9 := func(r, c int) int {
        maxVal := grid[r][c]
        for dr := 0; dr < 3; dr ++ {
            for dc := 0; dc < 3; dc ++ {
                val := grid[r+dr][c+dc]
                if val > maxVal {
                    maxVal = val
                }
            }
        }
        return maxVal
    }

    result := make([][]int, n-2)
    for i := 0; i < n-2; i++ {
        result[i] = make([]int, n-2)
    }

    for r := 0; r < n - 2; r ++ {
        for c := 0; c < n - 2; c ++ {
            result[r][c] = getMax9(r, c)
        }
    }
    return result
}