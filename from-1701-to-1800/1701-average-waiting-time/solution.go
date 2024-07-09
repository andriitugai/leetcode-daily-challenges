func averageWaitingTime(customers [][]int) float64 {
    waitingTotal := 0
    finishedTime := 0

    for _, c := range customers {
        arrivalTime, cookingTime := c[0], c[1]
        finishedTime = max(arrivalTime, finishedTime) + cookingTime
        waitingTotal += (finishedTime - arrivalTime)
    }
    return float64(waitingTotal)/float64(len(customers))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}