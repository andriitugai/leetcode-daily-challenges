func countMatchingSubarrays(nums []int, pattern []int) int {
    n := len(nums)
    model := make([]int, n - 1)
    for i := 0; i < n - 1; i ++ {
        if nums[i] > nums[i + 1] {
            model[i] = -1
        } else if nums[i] < nums[i + 1] {
            model[i] = 1
        }
    }

    m := len(pattern)
    result := 0
    for i := 0; i < n - m; i ++ {
        match := true
        for j := i; j < i + m; j ++ {
            if model[j] != pattern[j - i] {
                match = false
                break
            }
        }
        if match {
            result += 1
        }
    }
    return result
}