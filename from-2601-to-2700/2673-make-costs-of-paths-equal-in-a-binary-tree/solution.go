func minIncrements(n int, cost []int) int {
    result := 0
    var left, right int
    var mx, mn int
    for i := n - 1; i >= 0; i -- {
        if 2 * i + 2 < n {
            left = cost[2 * i + 1]
            right = cost[2 * i + 2]
            mx, mn = left, right
            if right > left {
                mx, mn = right, left
            }
            result += mx - mn
            cost[i] += mx
        }
    }
    return result
}