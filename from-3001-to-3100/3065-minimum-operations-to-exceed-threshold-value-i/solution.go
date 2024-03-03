func minOperations(nums []int, k int) int {
    result := 0
    for _, num := range nums {
        if num < k {
            result += 1
        }
    }
    return result
}