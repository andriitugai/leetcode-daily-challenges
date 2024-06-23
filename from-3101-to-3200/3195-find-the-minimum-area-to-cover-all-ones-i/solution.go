func minimumArea(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    var firstRow, lastRow, firstCol, lastCol int

    foundOne := false
    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            if grid[r][c] == 1 {
                foundOne = true
                firstRow = r
                break
            }
        }
        if foundOne {
            break
        }
    }

    foundOne = false
    for r := m - 1; r >= 0; r -- {
        for c := 0; c < n; c ++ {
            if grid[r][c] == 1 {
                foundOne = true
                lastRow = r
                break
            }
        }
        if foundOne {
            break
        }
    }

    foundOne = false 
    for c := 0; c < n; c ++ {
        for r := firstRow; r <= lastRow; r ++ {
            if grid[r][c] == 1 {
                foundOne = true
                firstCol = c
                break
            }
        }
        if foundOne {
            break
        }
    }

    foundOne = false 
    for c := n - 1; c >= 0; c -- {
        for r := firstRow; r <= lastRow; r ++ {
            if grid[r][c] == 1 {
                foundOne = true
                lastCol = c
                break
            }
        }
        if foundOne {
            break
        }
    }

    return (lastRow - firstRow + 1) * (lastCol - firstCol + 1)
}