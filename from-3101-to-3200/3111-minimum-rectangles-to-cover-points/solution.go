func minRectanglesToCoverPoints(points [][]int, w int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0]
    })
    result := 1
    x0 := points[0][0]
    for i := 1; i < len(points); i ++ {
        if points[i][0] > x0 + w {
            x0 = points[i][0]
            result += 1
        }
    }
    return result
}