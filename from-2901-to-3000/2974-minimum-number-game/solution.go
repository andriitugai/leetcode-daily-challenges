func numberGame(nums []int) []int {
    result := make([]int, len(nums))
    sort.Ints(nums)
    for i := 0; i < len(nums); i += 2 {
        result[i] = nums[i + 1]
        result[i + 1] = nums[i]
    }
    return result
}