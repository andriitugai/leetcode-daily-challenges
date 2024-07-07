func numberOfAlternatingGroups(colors []int) int {
    result := 0
    n := len(colors)
    prev := colors[n - 1]
    curr := colors[0]
    var next int
    for i := 0; i < n; i ++ {
        if i == n - 1 {
            next = colors[0]
        } else {
            next = colors[i + 1]
        }
        if curr != prev && curr != next {
            result += 1
        }
        prev = curr
        curr = next
    }
    return result
}