func getLastMoment(n int, left []int, right []int) int {
    time := 0
    for _, l := range left {
        if l > time {
            time = l
        }
    }
    for _, r := range right {
        if n - r > time {
            time = n - r
        }
    }
    return time
}