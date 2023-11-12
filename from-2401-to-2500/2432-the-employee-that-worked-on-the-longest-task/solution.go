unc hardestWorker(n int, logs [][]int) int {
    longestId := logs[0][0]
    longestTask := logs[0][1]

    for i := 1; i < len(logs); i++ {
        duration := logs[i][1] - logs[i-1][1]
        if duration > longestTask {
            longestTask = duration
            longestId = logs[i][0]
        } else if duration == longestTask && logs[i][0] < longestId {
            longestId = logs[i][0]
        }
    }
    return longestId
}