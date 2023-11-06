func digitCount(num string) bool {
    counts := make(map[int]int)
    for i := 0; i < len(num); i++ {
        counts[int(num[i] - '0')] += 1
    }

    for i := 0; i < len(num); i++ {
        if counts[i] != int(num[i] - '0') {
            return false
        }
    }
    return true
}