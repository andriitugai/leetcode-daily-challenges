func findMaxFish(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
            return 0
        }
        fish := grid[r][c]
        grid[r][c] = 0
        fish += dfs(r + 1, c)
        fish += dfs(r - 1, c)
        fish += dfs(r, c + 1)
        fish += dfs(r, c - 1)
        return fish
    }
    maxFish := 0
    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            if grid[r][c] > 0 {
                fish := dfs(r, c)
                if fish > maxFish {
                    maxFish = fish
                }
            }
        }
    }
    return maxFish
}