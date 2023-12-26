func numRollsToTarget(n int, k int, target int) int {
    cur := make([]int, target + 1)
    for i := 1; i <= k && i <= target; i ++ {
        cur[i] = 1
    }

    for x := 2; x <= n; x ++ {
        nxt := make([]int, target + 1)
        for i := x; i <= x * k && i <= target; i ++ {
            for j := i - 1; j >= i - k && j >= 0; j -- {
                nxt[i] = (nxt[i] + cur[j]) % 1000000007
            }
        }
        cur = nxt
    }

    return cur[target]
}

