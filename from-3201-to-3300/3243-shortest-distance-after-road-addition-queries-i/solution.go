func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    graph := make(map[int][]int)
    for i := 0; i < n - 1; i ++ {
        graph[i] = []int{i + 1}
    }

    result := []int{}
    for _, q := range queries {
        src, dst := q[0], q[1]
        graph[src] = append(graph[src], dst)
        result = append(result, bfs(graph, n - 1))
    }
    return result
}

func bfs(graph map[int][]int, tgt int) int {
    q := []int{0}
    visited := make(map[int]bool)
    steps := 0

    for len(q) > 0 {
        curLen := len(q)
        for i := 0; i < curLen; i ++ {
            city := q[0]
            if city == tgt {
                return steps
            }
            visited[city] = true
            q = q[1:]
            for _, dest := range graph[city] {
                if !visited[dest] {
                    q = append(q, dest)
                }
            }
        }
        steps += 1
    }
    return -1
}