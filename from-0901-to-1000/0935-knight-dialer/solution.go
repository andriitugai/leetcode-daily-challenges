func knightDialer(n int) int {
    if n == 1 {
        return 10
    }

    moves := map[int][]int {
        0: []int{4, 6},
        1: []int{8, 6},
        2: []int{7, 9},
        3: []int{4, 8},
        4: []int{0, 3, 9},
        5: []int{},
        6: []int{0, 1, 7},
        7: []int{2, 6},
        8: []int{1, 3},
        9: []int{2, 4},
    }
    mod := 1000000007

    memo := make([][]int, n)
    for i := 0; i < n; i ++ {
        memo[i] = make([]int, 10)
    }
    for i := 0; i < 10; i ++ {
        memo[0][i] = 1
    }

    for i := 1; i < n; i ++ {
        for j := 0; j < 10; j ++ {
            for _, d := range moves[j] {
                memo[i][j] += memo[i - 1][d]
            }
            memo[i][j] %= mod
        }
    }

    result := 0
    for i := 0; i < 10; i ++ {
        result += memo[n - 1][i]
        result %= mod
    }
    return result
}