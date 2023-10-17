func fullBloomFlowers(flowers [][]int, people []int) []int {
    bS := func(target int, arr []int) int {
        left, right := 0, len(arr) - 1
        for left <= right {
            mid := left + (right - left) / 2
            if arr[mid] <= target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        return left
    }

    result := make([]int, len(people))
    start := make([]int, len(flowers))
    end := make([]int, len(flowers))

    for i, flower := range flowers {
        start[i] = flower[0]
        end[i] = flower[1] + 1
    }

    sort.Ints(start)
    sort.Ints(end)
    for i, day := range people {
        started := bS(day, start)
        finished := bS(day, end)
        result[i] = started - finished
    }
    return result
}