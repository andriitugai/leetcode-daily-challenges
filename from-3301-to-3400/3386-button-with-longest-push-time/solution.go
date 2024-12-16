func buttonWithLongestTime(events [][]int) int {
    maxButton, maxTime := events[0][0], events[0][1]
    prevTime := events[0][1]

    for i := 1; i < len(events); i ++ {

        currButton, currTime := events[i][0], events[i][1]
        elapsedTime := currTime - prevTime
        if elapsedTime > maxTime {
            maxTime = elapsedTime
            maxButton = currButton
        } else if elapsedTime == maxTime && maxButton > currButton {
            maxButton = currButton
        }

        prevTime = currTime
    }
    return maxButton
}