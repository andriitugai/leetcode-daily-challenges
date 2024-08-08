func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
    y, x := rStart, cStart
    dY, dX := 0, 1
    twice := 2
    result := [][]int{}
    moves := 1
    nextMoves := 2

    for len(result) < rows * cols {
        if y >= 0 && y < rows && x >=0 && x < cols {
            result = append(result, []int{y, x})
        }
        y += dY
        x += dX
        moves -= 1
        if moves == 0 {
            dY, dX = dX, -dY
            twice -= 1
            if twice == 0 {
                twice = 2
                moves = nextMoves
                nextMoves += 1
            } else {
                moves = nextMoves - 1
            }
        }
    }
    return result
}