func minimumBoxes(apple []int, capacity []int) int {
    applesTotal := 0
    for _, app := range apple {
        applesTotal += app
    }
    sort.Slice(capacity, func(i, j int) bool {
        return capacity[j] < capacity[i]
    })
    boxes := 0
    for i := 0; i < len(capacity) && applesTotal > 0; i ++ {
        applesTotal -= capacity[i]
        boxes += 1
    }
    return boxes
}