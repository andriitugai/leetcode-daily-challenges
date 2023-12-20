func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    c := make(map[int]int)
    for i := 0; i < n; i ++ {
        for j := 0; j < n; j ++ {
            c[grid[i][j]] += 1
        }
    }

    var a, b int
    for k := 1; k <= n * n; k ++ {
        if c[k] == 2 {
            a = k
        } else if c[k] == 0 {
            b = k
        }
        if a != 0 && b != 0 {
            break
        }
    }
    return []int {a, b}
}