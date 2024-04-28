func canMakeSquare(grid [][]byte) bool {
    check := func(r0, c0 int) bool {
        cnt := 0
        for r := r0; r < r0 + 2; r ++ {
            for c := c0; c < c0 + 2; c ++ {
                if grid[r][c] == 'B' {
                    cnt += 1
                }
            }
        }
        return cnt != 2
    }

    for r := 0; r < 2; r ++ {
        for c := 0; c < 2; c ++ {
            if check(r, c) {
                return true
            }
        }
    }
    return false
}