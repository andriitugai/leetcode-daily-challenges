func maxScoreSightseeingPair(values []int) int {
    n := len(values)
    maxScore := 0
    maxIValue := values[0] + 0

    for i := 1; i < n; i ++ {
        score := maxIValue + values[i] - i
        if score > maxScore {
            maxScore = score
        }
        iValue := values[i] + i
        if iValue > maxIValue {
            maxIValue = iValue
        }
    }
    return maxScore
}