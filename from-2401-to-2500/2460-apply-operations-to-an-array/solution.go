func applyOperations(nums []int) []int {
    n := len(nums)
    for i := 0; i < n - 1; i ++ {
        if nums[i] == nums[i + 1] {
            nums[i] *= 2
            nums[i + 1] = 0
        }
    }
    result := make([]int, n)
    j := 0
    for i := 0; i < n; i ++ {
        if nums[i] > 0 {
            result[j] = nums[i]
            j += 1
        }
    }
    return result
}