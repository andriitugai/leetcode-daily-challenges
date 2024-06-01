func specialArray(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })
    i := 0
    for i < len(nums) && nums[i] > i {
        i += 1
    }
    if i < len(nums) && i == nums[i] {
        return -1
    }
    return i
}