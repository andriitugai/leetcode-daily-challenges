func findDuplicate(nums []int) int {
    counts := make(map[int]bool)
    for _, num := range nums {
        if counts[num] {
            return num
        }
        counts[num] = true
    }
    return -1
}