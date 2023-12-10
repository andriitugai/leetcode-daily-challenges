func transpose(matrix [][]int) [][]int {
    m, n := len(matrix), len(matrix[0])
    result := make([][]int, n)
    for i:=0; i < n; i++ {
        result[i] = make([]int, m)
    }
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            result[c][r] = matrix[r][c]
        }
    }
    return result
}