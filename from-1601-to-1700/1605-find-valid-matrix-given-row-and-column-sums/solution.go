func restoreMatrix(rowSum []int, colSum []int) [][]int {
    m, n := len(rowSum), len(colSum)
    result := make([][]int, m)
    
    for r := 0; r < m; r ++ {
        result[r] = make([]int, n)
        for c := 0; c < n; c ++ {
            result[r][c] = min(rowSum[r], colSum[c])
            rowSum[r] -= result[r][c]
            colSum[c] -= result[r][c]
        }
    }
    return result
}