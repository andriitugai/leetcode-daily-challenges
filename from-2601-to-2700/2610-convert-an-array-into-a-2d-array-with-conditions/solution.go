func findMatrix(nums []int) [][]int {
    counts := make(map[int]int)
    rows := 0
    for _, num := range(nums) {
        counts[num] += 1
        if counts[num] > rows {
            rows = counts[num]
        }
    }

    result := make([][]int, rows)
    
    for i := 0; i < rows; i ++ {
        result[i] = make([]int, 0)
    }

    for num, cnt := range counts {
        for i := 0; i < cnt; i ++ {
            result[i] = append(result[i], num)
        }
    }

    return result
}