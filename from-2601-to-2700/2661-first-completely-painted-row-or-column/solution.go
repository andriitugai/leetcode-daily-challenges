func firstCompleteIndex(arr []int, mat [][]int) int {
    mR, mC := make(map[int]int), make(map[int]int)
    m, n := len(mat), len(mat[0])
    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            mR[mat[r][c]] = r
            mC[mat[r][c]] = c
        }
    }
    rows := make([]int, m)
    cols := make([]int, n)
    var ri, ci int
    for i := 0; i < len(arr); i ++ {
        ri = mR[arr[i]]
        rows[ri] += 1
        if rows[ri] == n {
            return i
        }
        ci = mC[arr[i]]
        cols[ci] += 1
        if cols[ci] == m {
            return i
        }
    }
    return -1
}