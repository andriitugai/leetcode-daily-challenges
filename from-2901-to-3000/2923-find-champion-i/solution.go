func findChampion(grid [][]int) int {
    n := len(grid)
    for i := 0; i < n; i ++ {
        s := 0
        for j := 0; j < n; j ++ {
            s += grid[i][j]
        }
        if s == n - 1 {
            return i
        }
    }
    return -1
}