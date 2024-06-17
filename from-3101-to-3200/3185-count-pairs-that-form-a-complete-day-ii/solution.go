func countCompleteDayPairs(hours []int) int64 {
    var result int64 = 0
    reminders := make(map[int]int)
    for _, h := range hours {
        rem := (24 - h % 24) % 24
        if pairs, ok := reminders[rem]; ok {
            result += int64(pairs)
        }
        reminders[h % 24] += 1
    }
    return result
}