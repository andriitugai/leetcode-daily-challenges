func isArraySpecial(nums []int, queries [][]int) []bool {
    result := make([]bool, len(queries))
    pairs := make([]int, len(nums) - 1)
    for i := 1; i < len(nums); i ++ {
        pairs[i - 1] = (nums[i - 1] % 2) ^ (nums[i] % 2)
    }
    preSum := make([]int, len(nums))
    preSum[0] = 1
    for i, val := range pairs {
        preSum[i+1] = preSum[i] + val
    }

    for i, q := range queries {
        if preSum[q[1]] - preSum[q[0]] == q[1] - q[0] {
            result[i] = true
        }
    }
    return result
}