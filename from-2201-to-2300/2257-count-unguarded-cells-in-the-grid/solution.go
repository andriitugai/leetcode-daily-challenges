func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    grid := make([][]int, m)
    for i := 0; i < m; i ++ {
        grid[i] = make([]int, n)
    }
    // 0 - free, 1 - guard, 2 - wall, 3 - guardable

    for _, g := range guards {
        grid[g[0]][g[1]] = 1
    }
    for _, w := range walls {
        grid[w[0]][w[1]] = 2
    }

    for _, g := range guards {
        r, c := g[0], g[1]
        for i := r - 1; i >= 0; i -- {
            if grid[i][c] == 1 || grid[i][c] == 2 {
                break
            }
            grid[i][c] = 3
        }
        for i := r + 1; i < m; i ++ {
            if grid[i][c] == 1 || grid[i][c] == 2 {
                break
            }
            grid[i][c] = 3
        }
        for j := c - 1; j >= 0; j -- {
            if grid[r][j] == 1 || grid[r][j] == 2 {
                break
            }
            grid[r][j] = 3
        }
        for j := c + 1; j < n; j ++ {
            if grid[r][j] == 1 || grid[r][j] == 2 {
                break
            }
            grid[r][j] = 3
        }
    }

    result := 0
    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            if grid[r][c] == 0 {
                result += 1
            }
        }
    }
    return result
}