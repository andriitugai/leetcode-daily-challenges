func countDays(days int, meetings [][]int) int {
    sort.Slice(meetings, func(i, j int) bool{
        return meetings[i][0] < meetings[j][0]
    })
    n := len(meetings)
    result := abs(meetings[0][0] - 1)
    for i := 1; i < n; i ++ {
        if meetings[i][0] <= meetings[i - 1][1] {
            if meetings[i][1] < meetings[i - 1][1] {
                meetings[i][1] = meetings[i - 1][1]
            }
        } else {
            dy := meetings[i][0] - meetings[i - 1][1]
            result += dy - 1
        }
    }
    result += abs(meetings[n - 1][1] - days)
    return result
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}