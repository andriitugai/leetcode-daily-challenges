func buildArray(target []int, n int) []string {
    result := []string{}
    targetMap := make(map[int]bool)
    for _, tgt := range target {
        targetMap[tgt] = true
    }
    goal := target[len(target) - 1]

    for nxt := 1; nxt <= n; nxt++ {
        result = append(result, "Push")
        if !targetMap[nxt] {
            result = append(result, "Pop")
        }
        if nxt == goal {
            break
        }
    }
    return result
}