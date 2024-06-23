func minOperations(nums []int) int {
    result := 0
    n := len(nums)
    for i := 0; i < n - 2; i ++ {
        if nums[i] == 0 {
            result += 1
            flip(&nums[i + 1])
            flip(&nums[i + 2])
        }
    }
    if nums[n - 2] == 0 || nums[n - 1] == 0 {
        return -1
    }
    return result
}

func flip(a *int) {
    *a = (*a + 1) % 2
}