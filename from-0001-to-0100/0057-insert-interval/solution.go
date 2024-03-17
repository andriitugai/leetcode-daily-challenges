func insert(intervals [][]int, newInterval []int) [][]int {
    result := make([][]int, 0) 
    start, end := newInterval[0], newInterval[1]

    for _, interval := range intervals {
        s, e := interval[0], interval[1]
        if s > end {
            result = append(result, []int{start, end})
            start, end = s, e
        } else if e < start {
            result = append(result, interval)
        } else {
            if s < start {
                start = s
            }
            if e > end {
                end = e
            }
        }
    }
    result = append(result, []int{start, end})
    return result
}