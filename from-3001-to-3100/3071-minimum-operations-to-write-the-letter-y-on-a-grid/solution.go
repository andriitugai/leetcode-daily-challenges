func minimumOperationsToWriteY(grid [][]int) int {
    n := len(grid)
    onArea := n + n / 2
    outArea := n * n - onArea
    mid := n / 2

    onY := func(r, c int) bool {
        if r <= mid && (r == c || r == n - 1 - c) || r > mid && c == mid {
            return true
        } 
        return false
    }

    onCells := make(map[int]int)
    outCells := make(map[int]int)
    for r := 0; r < n; r ++ {
        for c := 0; c < n; c ++ {
            if onY(r, c) {
                onCells[grid[r][c]] += 1
            } else {
                outCells[grid[r][c]] += 1
            }
        }
    }

    y0r1 := (onArea - onCells[0]) + (outArea - outCells[1])
    y0r2 := (onArea - onCells[0]) + (outArea - outCells[2])
    y1r0 := (onArea - onCells[1]) + (outArea - outCells[0])
    y1r2 := (onArea - onCells[1]) + (outArea - outCells[2])
    y2r0 := (onArea - onCells[2]) + (outArea - outCells[0])
    y2r1 := (onArea - onCells[2]) + (outArea - outCells[1])
    result := y0r1
    if y0r2 < result { result = y0r2 }
    if y1r0 < result { result = y1r0 }
    if y1r2 < result { result = y1r2 }
    if y2r0 < result { result = y2r0 }
    if y2r1 < result { result = y2r1 }
    return result
}