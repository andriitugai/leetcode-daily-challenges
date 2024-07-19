func luckyNumbers (matrix [][]int) []int {
    result := []int{}
    m, n := len(matrix), len(matrix[0])

    for r := 0; r < m; r ++ {
        minVal := matrix[r][0]
        minCol := 0
        for c := 1; c < n; c ++ {
            if matrix[r][c] < minVal {
                minVal = matrix[r][c]
                minCol = c
            }
        }
        lucky := true
        for row := 0; row < m; row ++ {
            if matrix[row][minCol] > minVal {
                lucky = false
                break
            }
        }
        if lucky {
            result = append(result, minVal)
        }
    }
    return result
}