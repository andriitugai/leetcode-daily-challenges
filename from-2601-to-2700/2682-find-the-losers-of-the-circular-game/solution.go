func circularGameLosers(n int, k int) []int {
    players := make([]bool, n)
    curr := 0
    for round := 1; !players[curr]; round ++ {
        players[curr] = true
        curr = (curr + round * k) % n
    }

    result := []int{}
    for i := 0; i < n; i ++ {
        if !players[i] {
            result = append(result, i + 1)
        }
    }
    return result
}