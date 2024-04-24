func maxRotateFunction(nums []int) int {
    total := 0
    n := len(nums)
    curr, prev := 0, 0
    for i := 0; i < n; i ++ {
        total += nums[i]
        prev += i * nums[i]
    }
    result := prev
    for i := 1; i < n; i ++ {
        curr = prev + total - n * nums[n - i]
        if curr > result {
            result = curr
        }
        prev = curr
    }
    return result
}