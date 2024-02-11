type State struct {
    row int
    col1 int
    col2 int
}

var dirs []int = []int{-1, 0, 1}

func cherryPickup(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    memo := make(map[State]int)

    var dp func(r , c1, c2 int) int
    dp = func(r, c1, c2 int) int {
        if r == m || c1 < 0 || c1 == n || c2 < 0 || c2 == n {
            return 0
        }
        curr := State{row: r, col1: c1, col2: c2}
        if res, ok := memo[curr]; ok {
            return res
        }
        cherries := grid[r][c1]
        if c1 != c2 {
            cherries += grid[r][c2]
        }

        hvst := 0
        for _, dc1 := range dirs {
            for _, dc2 := range dirs {
                dh := dp(r + 1, c1 + dc1, c2 + dc2) 
                if dh > hvst {
                    hvst = dh
                }
            }
        }

        result := cherries + hvst
        memo[curr] = result
        return result
    }
    return dp(0, 0, n - 1)
}