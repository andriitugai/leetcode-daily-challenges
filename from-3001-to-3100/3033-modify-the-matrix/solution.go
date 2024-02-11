func modifiedMatrix(matrix [][]int) [][]int {
    m, n := len(matrix), len(matrix[0])
    maxCols := make([]int, n)
    copy(maxCols, matrix[0])

    for i := 1; i < m; i ++ {
        for j := 0; j < n; j ++ {
            if maxCols[j] < matrix[i][j] {
                maxCols[j] = matrix[i][j]
            }
        }
    }

    answer := make([][]int, m)
    for i := 0; i < m; i ++ {
        row := make([]int, n)
        copy(row, matrix[i])

        for j := 0; j < n; j ++ {
            if row[j] == -1 {
                row[j] = maxCols[j]
            }
        }
        answer[i] = row
    }
    return answer
}