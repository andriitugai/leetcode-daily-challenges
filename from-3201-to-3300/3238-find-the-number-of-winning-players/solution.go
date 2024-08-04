func winningPlayerCount(n int, pick [][]int) int {
    counts := make([]map[int]int, n)
    for i := 0; i < n; i ++ {
        counts[i] = make(map[int]int)
    }
    for _, p := range pick {
        player, color := p[0], p[1]
        counts[player][color] += 1
    }
    result := 0
    for i := 0; i < n; i ++ {
        for _, v := range counts[i] {
            if v > i {
                result += 1
                break
            }
        }
    }
    return result
}