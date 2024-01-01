func hasTrailingZeros(nums []int) bool {
    count := 0
    for _, num := range nums {
        if num % 2 == 0 {
            count += 1
        }
    }
    return count > 1
}