func areaOfMaxDiagonal(dimensions [][]int) int {
    maxDiagSquared := 0
    maxArea := 0
    for i := 0; i < len(dimensions); i ++ {
        diagSquared := dimensions[i][0] * dimensions[i][0] + dimensions[i][1] * dimensions[i][1]
        area := dimensions[i][0] * dimensions[i][1]
        if diagSquared > maxDiagSquared {
            maxDiagSquared = diagSquared
            maxArea = area
        } else if diagSquared == maxDiagSquared && area > maxArea {
            maxArea = area
        }
    }
    return maxArea
}