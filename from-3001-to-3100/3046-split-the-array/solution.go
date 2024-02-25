func isPossibleToSplit(nums []int) bool {
    counts := make(map[int]int)
    for _, num := range nums {
        counts[num] += 1
        if counts[num] > 2 {
            return false
        }
    }
    return true
}