func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
    marbles := make(map[int]int)
    for _, pos := range nums {
        marbles[pos] += 1
    }

    for i := 0; i < len(moveFrom); i ++ {
        if moveFrom[i] != moveTo[i] {
            marbles[moveTo[i]] += marbles[moveFrom[i]]
            marbles[moveFrom[i]] = 0
        }
    }
    result := []int{}
    for pos, val := range marbles {
        if val > 0 {
            result = append(result, pos)
        }
    }
    sort.Ints(result)
    return result
}