func minOperations(nums []int) int {
    prev := nums[0]
    result := 0
    for i := 1; i < len(nums); i ++ {
        if nums[i] > prev {
            prev = nums[i]
        } else {
            prev = prev + 1
            result += (prev - nums[i])
        }
    }
    return result
}