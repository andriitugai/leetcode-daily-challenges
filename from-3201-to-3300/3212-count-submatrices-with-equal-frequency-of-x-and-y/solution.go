func numberOfSubmatrices(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    preCol, preRow := make([]int, n + 1), make([]int, n + 1)
    result := 0
    for i := 0; i < m; i ++ {
        currX, currY := 0, 0
        for j := 0; j < n; j ++ {
            if grid[i][j] == 'X' {
                currX += 1
            } else if grid[i][j] == 'Y' {
                currY += 1
            }
            preCol[j + 1] += currX
            preRow[j + 1] += currY

            if preCol[j + 1] == preRow[j + 1] &&preCol[j + 1] > 0 {
                result += 1
            }  
        }
    }
    return result
}