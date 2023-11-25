func findNonMinOrMax(nums []int) int {
    minVal, maxVal := nums[0], nums[0]
    for i := 1; i < len(nums); i ++ {
        if nums[i] < minVal {
            minVal = nums[i]
        } else if nums[i] > maxVal {
            maxVal = nums[i]
        }
    }

    for i := 0; i < len(nums); i ++ {
        if nums[i] != minVal && nums[i] != maxVal {
            return nums[i]
        }
    }
    return -1
}