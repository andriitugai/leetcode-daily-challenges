func findSubarrays(nums []int) bool {
    sums := make(map[int]bool)
    for i := 0; i < len(nums) - 1; i ++ {
        s := nums[i] + nums[i+1]
        if sums[s] {
            return true
        }
        sums[s] = true
    }
    return false
}