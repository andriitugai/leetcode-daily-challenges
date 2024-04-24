func differenceOfDistinctValues(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    var tl, br map[int]bool
    var ntl, nbr int

    answer := make([][]int, m)
    for i := 0; i < m; i ++ {
        answer[i] = make([]int, n)
    }

    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            tl = make(map[int]bool)
            for i, j := r - 1, c - 1; i >= 0 && j >= 0; i, j = i - 1, j - 1 {
                tl[grid[i][j]] = true
            }
            ntl = len(tl)
            br = make(map[int]bool)
            for i, j := r + 1, c + 1; i < m && j < n; i, j = i + 1, j + 1 {
                br[grid[i][j]] = true
            }
            nbr = len(br)
            if nbr > ntl {
                answer[r][c] = nbr - ntl
            } else {
                answer[r][c] = ntl - nbr
            }
        }
    }
    return answer
}