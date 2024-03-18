func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][1] < points[j][1]
    })
    arrowX := points[0][1]
    count := 1
    for _, p := range points {
        s, e := p[0], p[1]
        if s > arrowX {
            count += 1
            arrowX = e
        }
    }
    return count
}