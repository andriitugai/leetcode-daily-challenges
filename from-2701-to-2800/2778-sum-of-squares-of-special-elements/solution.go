func sumOfSquares(nums []int) int {
    result := 0
    n := len(nums)
    for i := 1; i <= n; i ++ {
        if n % i == 0 {
            result += nums[i-1] * nums[i-1]
        }
    }
    return result
}