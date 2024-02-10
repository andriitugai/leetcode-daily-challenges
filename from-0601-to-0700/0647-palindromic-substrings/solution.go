func countSubstrings(s string) int {
    n := len(s)
    if n < 2 {
        return n
    }

    count := n
    dp := make([][]int, n)
    for i := 0; i < n; i ++ {
        dp[i] = make([]int, n)
        dp[i][i] = 1
    }

    for col := 1; col < n; col ++ {
        for row := 0; row < col; row ++ {
            if s[row] == s[col]{
                if row == col - 1 || dp[row + 1][col - 1] == 1{
                    dp[row][col] = 1
                    count += 1
                }
            }
        }
    } 
    return count
}