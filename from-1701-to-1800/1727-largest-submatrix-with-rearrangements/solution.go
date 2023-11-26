func largestSubmatrix(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])

    for i := 1; i < m; i ++ {
        for j := 0; j < n; j ++ {
            if matrix[i][j] == 1 {
                matrix[i][j] += matrix[i - 1][j]
            }
        }
    }
    result := 0
    for _, row := range matrix {
        sort.Slice(row, func(i, j int) bool {
            return row[i] > row[j]
        })
        for j := 0; j < n; j ++ {
            if result < row[j] * (j + 1) {
                result = row[j] * (j + 1)
            }
        }
    }
    return result
}