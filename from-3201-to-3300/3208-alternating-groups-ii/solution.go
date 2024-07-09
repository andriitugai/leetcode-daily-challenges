func numberOfAlternatingGroups(colors []int, k int) int {
    n := len(colors)
    alt := 1
    result := 0
    
    for i := 1; i < n + k - 1; i ++ {
        if colors[i % n] == colors[(n + i - 1) % n] {
            alt = 1
        } else {
            alt += 1
        }
        if alt >= k {
            result += 1
        }
    }
    return result
}