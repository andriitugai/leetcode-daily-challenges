func countCompleteDayPairs(hours []int) int {
    result := 0
    reminders := make(map[int]int)
    for _, h := range hours {
        rem := (24 - h % 24) % 24
        if pairs, ok := reminders[rem]; ok {
            result += pairs
        }
        reminders[h % 24] += 1
    }
    return result
}