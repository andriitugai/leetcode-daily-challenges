func transformArray(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    ptr := n - 1
    for i := 0; i < n; i ++ {
        if nums[i] % 2 == 1 {
            result[ptr] = 1
            ptr -= 1
        }
    }
    return result
}