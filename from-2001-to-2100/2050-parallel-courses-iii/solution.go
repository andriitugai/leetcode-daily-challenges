func minimumTime(n int, relations [][]int, time []int) int {
    outcount := make([]int, n + 1)
    graph := make(map[int][]int)
    for _, rel := range relations {
        graph[rel[1]] = append(graph[rel[1]], rel[0])
        outcount[rel[0]] += 1
    }
    dp := make([]int, n + 1)
    for i := 0; i < n + 1; i++ {
        dp[i] = -1
    }
    
    month := 0
    for i := 1; i < n + 1; i++ {
        if outcount[i] == 0 {
            month = max(month, dfs(i, graph, time, dp))
        }
    }
    return month
}

func dfs(course int, graph map[int][]int, time []int, dp []int) int { // returns time for ```course``` course
    if dp[course] != -1 {
        return dp[course]
    }

    self := time[course - 1]
    maxDur := 0
    for _, child := range graph[course] {
        maxDur = max(maxDur, dfs(child, graph, time, dp))
    }
    dp[course] = self + maxDur
    return dp[course]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
    