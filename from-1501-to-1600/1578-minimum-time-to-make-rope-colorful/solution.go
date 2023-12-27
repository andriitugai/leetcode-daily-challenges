func minCost(colors string, neededTime []int) int {
    result := 0
    prevColor := colors[0]
    prevIdx := 0
    for i := 1; i < len(colors); i ++ {
        if colors[i] != prevColor {
            prevColor = colors[i]
            prevIdx = i
            continue
        }
        if neededTime[i] < neededTime[prevIdx] {
            result += neededTime[i]
        } else {
            result += neededTime[prevIdx]
            prevIdx = i
        }
    }
    return result
}