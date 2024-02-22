func findJudge(n int, trust [][]int) int {
    ppl := make([]int, n + 1)
    for i := 0; i < len(trust); i ++ {
        ppl[trust[i][1]] += 1
        ppl[trust[i][0]] -= 1
    }

    candidates := 0
    judge := -1
    for i := 1; i < n + 1; i ++ {
        if ppl[i] == n - 1 {
            candidates += 1
            if candidates > 1 {
                return -1
            }
            judge = i
        }
    }
    return judge
}