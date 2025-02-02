func check(nums []int) bool {
    rotationPoints := 0
    n := len(nums)
    for i := 1; i < n; i ++ {
        if nums[i - 1] > nums[i] {
            rotationPoints += 1
        }
    }
    if nums[n - 1] > nums[0] {
        rotationPoints += 1
    }
    return rotationPoints <= 1
}