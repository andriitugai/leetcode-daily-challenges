const MOD int = 1000000007
var DIRS [][2]int = [][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    memo := make([][][]int, m)
    for i := 0; i < m; i++ {
        memo[i] = make([][]int, n)
        for j := 0; j < n; j ++ {
            memo[i][j] = make([]int, maxMove + 1)
            for k := 0; k <= maxMove; k ++ {
                memo[i][j][k] = -1
            }
        }
    }

    var dfs func(curMoves, curRow, curCol int) int
    dfs = func(curMoves, curRow, curCol int) int {
        if curRow < 0 || curRow == m || curCol < 0 || curCol == n {
            return 1
        }
        if curMoves == 0 {
            return 0
        }
        if memo[curRow][curCol][curMoves] >= 0 {
            return memo[curRow][curCol][curMoves]
        }
        result := 0
        for _, dir := range DIRS {
            result += dfs(curMoves - 1, curRow + dir[0], curCol + dir[1]) % MOD
        }
        result %= MOD
        memo[curRow][curCol][curMoves] = result
        return result
    }
    return dfs(maxMove, startRow, startColumn)
}