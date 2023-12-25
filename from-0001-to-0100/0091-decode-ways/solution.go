func dfs(i int, s string, dp map[int]int) int {
    v, ok := dp[i]
    if ok {
        return v
    }
    if s[i] == '0' {
        return 0
    }
    v = dfs(i + 1, s, dp)
    if i + 1 < len(s) && (s[i] == '1' || s[i] == '2' && strings.Contains("0123456", string(s[i+1]))) {
        v += dfs(i + 2, s, dp)
    }
    dp[i] = v
    return v
}

func numDecodings(s string) int {
    dp := map[int]int {
        len(s): 1,
    }
    return dfs(0, s, dp)
}
