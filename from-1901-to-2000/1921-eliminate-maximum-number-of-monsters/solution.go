func eliminateMaximum(dist []int, speed []int) int {
    n := len(dist)
    time := make([]float64, n)

    for i := 0; i < n; i ++ {
        time[i] = float64(dist[i]) / float64(speed[i])
    }

    sort.Float64s(time)

    for i := 0; i < n; i ++ {
        if time[i] <= float64(i) {
            return i
        }
    }
    return n
}