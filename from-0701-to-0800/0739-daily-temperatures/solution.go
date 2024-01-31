func dailyTemperatures(temperatures []int) []int {
    result := make([]int, len(temperatures))
    q := []int{}

    for idx, currTemp := range temperatures {
        for len(q) > 0 && currTemp > temperatures[q[len(q) - 1]] {
            result[q[len(q) - 1]] = idx - q[len(q) - 1]
            q = q[:len(q) - 1]
        }
        q = append(q, idx)
    }

    return result
}