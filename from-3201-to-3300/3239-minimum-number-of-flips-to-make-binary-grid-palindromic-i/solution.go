func minFlips(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    rowFlips := 0
    for r := 0; r < m; r ++ {
        left, right := 0, n - 1
        for left < right {
            if grid[r][left] != grid[r][right] {
                rowFlips += 1
            }
            left += 1
            right -= 1
        }
    }

    colFlips := 0
    for c := 0; c < n; c ++ {
        top, bottom := 0, m - 1
        for top < bottom {
            if grid[top][c] != grid[bottom][c] {
                colFlips += 1
            }
            top += 1
            bottom -= 1
        }
    }

    if colFlips < rowFlips {
        return colFlips
    }
    return rowFlips
}