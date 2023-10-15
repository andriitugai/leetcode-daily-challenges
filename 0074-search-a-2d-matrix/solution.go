func searchMatrix(matrix [][]int, target int) bool {
    var m, n int = len(matrix), len(matrix[0])
    var i, j int = m - 1, 0

    for i >= 0 && j < n {
        if target == matrix[i][j] {
            return true
        }
        if target < matrix[i][j] {
            i--
        } else {
            j++
        }
    }
    return false
}