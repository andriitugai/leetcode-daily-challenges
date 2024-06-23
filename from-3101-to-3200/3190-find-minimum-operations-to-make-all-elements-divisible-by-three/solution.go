func minimumOperations(nums []int) int {
    result := 0
    for _, num := range nums {
        if num % 3 != 0 {
            result += 1
        }
    }
    return result
}