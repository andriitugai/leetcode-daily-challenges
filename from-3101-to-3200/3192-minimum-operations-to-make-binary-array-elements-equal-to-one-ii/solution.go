func minOperations(nums []int) int {
    result := 0
    flipped := false
    for _, num := range nums {
        if num == 0 && !flipped || num == 1 && flipped {
            result += 1
            flipped = !flipped
        }
    }
    return result
}