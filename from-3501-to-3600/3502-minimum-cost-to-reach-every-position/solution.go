func minCosts(cost []int) []int {
    currMin := cost[0]
    n := len(cost)
    result := make([]int, n)
    for i := 0; i < n; i ++ {
        result[i] = min(currMin, cost[i])
        currMin = result[i]
    }
    return result
}