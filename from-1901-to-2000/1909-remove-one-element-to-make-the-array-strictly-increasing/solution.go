func canBeIncreasing(nums []int) bool {
    used := false
    for i := 1; i < len(nums); i ++ {
        if nums[i] <= nums[i - 1] {
            if used {
                return false
            }
            if i == 1 || nums[i] > nums[i-2] {
                nums[i-1] = nums[i]
            } else {
                nums[i] = nums[i - 1]
            }
            used = true
        }
    }
    return true
}