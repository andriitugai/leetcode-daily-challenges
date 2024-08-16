func maxDistance(arrays [][]int) int {
    minVal := arrays[0][0]
    maxVal := arrays[0][len(arrays[0]) - 1]
    maxDist := 0

    for _, arr := range arrays[1:] {
        first, last := arr[0], arr[len(arr) - 1]
        maxDist = max(maxDist, max(abs(last - minVal), abs(first - maxVal)))
        minVal = min(minVal, first)
        maxVal = max(maxVal, last)
    }

    return maxDist  
}


func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}