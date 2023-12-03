func findCenter(edges [][]int) int {
    e := map[int]bool{}
    for i := 0; i < 2; i ++ {
        v := edges[i][0]
        if e[v] {
            return v
        }
        e[v] = true
        v = edges[i][1]
        if e[v] {
            return v
        }
        e[v] = true
    }
    return -1
}