func maxWidthOfVerticalArea(points [][]int) int {
    sort.Slice(points, func(i, j int) bool{
        return points[i][0] < points[j][0]
    })
    maxWidth := 0
    for i := 1; i < len(points); i ++ {
        width := points[i][0] - points[i - 1][0]
        if width > maxWidth {
            maxWidth = width
        }
    }
    return maxWidth
}