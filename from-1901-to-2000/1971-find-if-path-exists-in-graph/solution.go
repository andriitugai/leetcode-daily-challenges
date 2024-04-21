func validPath(n int, edges [][]int, source int, destination int) bool {
    adj := make(map[int][]int)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], e[1])
        adj[e[1]] = append(adj[e[1]], e[0])
    }
    seen := make(map[int]bool)
    q := []int{source}

    for len(q) > 0 {
        v := q[0]
        if v == destination {
            return true
        }
        q = q[1:]
        if seen[v] {
            continue
        }
        seen[v] = true
        for _, a := range adj[v] {
            q = append(q, a)
        }
    }
    return false
}