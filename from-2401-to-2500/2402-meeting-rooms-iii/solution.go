func mostBooked(n int, meetings [][]int) int {
    inf := math.Inf(1)
    avail := make([]int, n)
    mCount := make([]int, n)
    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][0] < meetings[j][0]
    })
    for i := 0; i < len(meetings); i ++ {
        start, end := meetings[i][0], meetings[i][1]
        minRoomAvailTime := inf
        minAvailTimeRoom := 0
        foundUnused := false
        for i := 0; i < n; i ++ {
            if avail[i] <= start {
                foundUnused = true
                mCount[i] += 1
                avail[i] = end
                break
            }
            if minRoomAvailTime > float64(avail[i]) {
                minRoomAvailTime = float64(avail[i])
                minAvailTimeRoom = i
            }
        }
        if !foundUnused {
            avail[minAvailTimeRoom] += end - start
            mCount[minAvailTimeRoom] += 1
        }
    }
    maxMeetings := 0
    maxMeetingIndex := -1
    for i := 0; i < n; i ++ {
        if mCount[i] > maxMeetings {
            maxMeetings = mCount[i]
            maxMeetingIndex = i
        }
    }
    return maxMeetingIndex
}