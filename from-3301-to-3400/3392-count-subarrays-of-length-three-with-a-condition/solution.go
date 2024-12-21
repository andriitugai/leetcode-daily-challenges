func countSubarrays(nums []int) int {
    result := 0
    for i := 1; i < len(nums) - 1; i ++ {
        if (nums[i-1] + nums[i+1]) * 2 == nums[i] {
            result += 1
        }
    }
    return result
}