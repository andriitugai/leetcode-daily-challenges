func minOperations(nums []int, k int) int {
    r := k
    for _, num := range nums {
        r ^= num
    }
    result := 0
    for r > 0 {
        r = r & (r - 1)
        result += 1
    }
    return result
}