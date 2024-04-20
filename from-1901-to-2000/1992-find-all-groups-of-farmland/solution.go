func findFarmland(land [][]int) [][]int {
    result := [][]int{}
    m, n := len(land), len(land[0])

    findRB := func(r0, c0 int) {
        r1, c1 := r0, c0
        for r1 < m && land[r1][c0] == 1 {
            r1 += 1
        }
        for c1 < n && land[r0][c1] == 1 {
            c1 += 1
        }
        for r := r0; r < r1; r ++ {
            for c := c0; c < c1; c ++ {
                land[r][c] = 0
            }
        }
        result = append(result, []int{r0, c0, r1-1, c1-1})
    }
    for i := 0; i < m; i ++ {
        for j := 0; j < n; j ++ {
            if land[i][j] == 1 {
                findRB(i, j)
            }
        }
    }
    return result
}