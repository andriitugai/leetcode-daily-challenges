func numSpecial(mat [][]int) int {
    m, n := len(mat), len(mat[0])

    checkRow := func(row int) (bool, int) {
        specCol := -1
        for c := 0; c < n; c++ {
            if mat[row][c] == 1 {
                if specCol < 0 {
                    specCol = c
                } else {
                    return false, -1
                }
            }
        }
        result := true
        if specCol < 0 {
            result = false
        }
        return result, specCol
    }
    checkCol := func(col int) (bool, int) {
        res := -1
        for row := 0; row < m; row ++ {
            if mat[row][col] == 1 {
                if res < 0 {
                    res = row
                } else {
                    return false, -1
                }
            }
        }
        return true, res
    }
    result := 0
    for row := 0; row < m; row ++ {
        isSpecRow, colNum := checkRow(row)
        if isSpecRow {
            if isSpecCol, _ := checkCol(colNum); isSpecCol {
                result += 1
            }
        }
    } 
    return result
}