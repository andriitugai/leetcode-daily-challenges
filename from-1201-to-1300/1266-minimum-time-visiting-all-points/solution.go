func minTimeToVisitAllPoints(points [][]int) int {
    abs := func(a int) int {
        if a < 0 {
            return -a
        }
        return a
    }

    max := func(a, b int) int {
        if b > a {
            return b
        }
        return a
    }

    result := 0
    curr := points[0]
    for i := 1; i < len(points); i++ {
        dx := abs(curr[0] - points[i][0])
        dy := abs(curr[1] - points[i][1])
        result += max(dx, dy)
        curr = points[i]
    }
    return result
}