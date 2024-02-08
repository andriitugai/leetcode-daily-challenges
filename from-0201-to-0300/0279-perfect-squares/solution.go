func numSquares(n int) int {
    squares := []int{0}
    var sq int = 0
    for i := 1; sq <= n; i ++ {
        sq = i * i
        squares = append(squares, sq)
    }

    // BFS traversal
    q := make([]int, 0)
    q = append(q, n)
    level := 0
    visited := make(map[int]bool)

    for len(q) > 0 {
        size := len(q)
        level += 1
        for i := 0; i < size; i++ {
            num := q[0]
            q = q[1:]
            for _, square := range squares {
                next := num - square
                if next < 0 {
                    break
                }
                if next == 0 {
                    return level
                }
                if visited[next] {
                    continue
                }
                visited[next] = true
                q = append(q, next)
            }
        }
    }
    return -1 
}