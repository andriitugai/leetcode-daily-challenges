func firstMissingPositive(nums []int) int {
    n := len(nums)
    i := 0
    for i < n {
        c := nums[i] - 1
        if nums[i] > 0 && c < n {
            if nums[i] != nums[c] {
                nums[i], nums[c] = nums[c], nums[i]
            } else {
                i += 1
            }
        } else {
            i += 1
        }
    }

    for i := 0; i < n; i++ {
        if nums[i] != i + 1 {
            return i + 1
        }
    }
    return n + 1
}