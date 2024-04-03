func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    if len(word) > m * n {
        return false
    }
    dirs := []int{1, 0, -1, 0, 1}

    var dfs func(r, c, idx int) bool
    dfs = func(r, c, idx int) bool {
        if idx == len(word) {
            return true
        }
        if r < 0 || r >= m || c < 0 || c >= n {
            return false
        }
        ch := board[r][c]
        if ch != word[idx] {
            return false
        }
        board[r][c] = '*'
        for i := 0; i < 4; i ++ {
            if dfs(r + dirs[i], c + dirs[i+1], idx + 1) {
                return true
            }
        }
        board[r][c] = ch
        return false
    }
    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            if board[r][c] == word[0] && dfs(r, c, 0) {
                return true
            }
        }
    }
    return false
}