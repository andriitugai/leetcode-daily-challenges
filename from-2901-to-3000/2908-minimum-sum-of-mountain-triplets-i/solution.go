func minimumSum(nums []int) int {
    minSum := 2000
    n := len(nums)
    for i := 0; i < n - 2; i ++ {
        for j := i + 1; j < n - 1; j ++ {
            for k := j + 1; k < n; k ++ {
                if nums[i] < nums[j] && nums[j] > nums[k] && nums[i] + nums[j] + nums[k] < minSum {
                    minSum = nums[i] + nums[j] + nums[k]
                }
            }
        }
    }
    if minSum == 2000 {
        return -1
    }
    return minSum
}