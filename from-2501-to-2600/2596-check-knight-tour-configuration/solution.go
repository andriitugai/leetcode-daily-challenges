func checkValidGrid(grid [][]int) bool {
    n := len(grid)
    deltas := [][]int {
        []int{1, 2}, []int{2, 1}, []int{2, -1}, []int{1, -2},
        []int{-1, -2}, []int{-2, -1}, []int{-2, 1}, []int{-1, 2},
    }
    getNextPos := func(currPos []int, nextVal int) ([]int, error) {
        cr, cc := currPos[0], currPos[1]
        for _, delta := range deltas {
            nr, nc := cr + delta[0], cc + delta[1]
            if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] == nextVal {
                return []int{nr, nc}, nil
            }
        }
        return []int{-1, -1}, errors.New("No such value")
    }
    currPos := []int{0, 0}
    found := false
    for r := 0; r < n; r ++ {
        for c := 0; c < n; c ++ {
            if grid[r][c] == 0 {
                currPos[0] = r
                currPos[1] = c
                found = true
                break
            }
        }
        if found {
            break
        }
    }
    if currPos[0] != 0 || currPos[1] != 0 {
        return false
    }

    for v := 1; v < n * n; v ++ {
        if nextPos, err := getNextPos(currPos, v); err != nil {
            return false
        } else {
            currPos = nextPos
        }
    }
    return true
}