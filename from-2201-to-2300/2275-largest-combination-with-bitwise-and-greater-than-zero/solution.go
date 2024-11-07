func largestCombination(candidates []int) int {
    result := 0
    bitmask := 1
    for i := 0; i < 32; i ++ {
        count := 0
        for _, n := range(candidates) {
            if bitmask & n != 0 {
                count += 1
            }
        }
        if count > result {
            result = count
        }
        bitmask <<= 1
    }
    return result
}