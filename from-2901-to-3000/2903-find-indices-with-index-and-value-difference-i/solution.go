func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    n := len(nums)
    abs := func(a int) int{
        if a < 0 {
            return -a
        }
        return a
    }
    for i := 0; i < n - indexDifference; i ++ {
        for j := i + indexDifference; j < n; j ++ {
            if abs(nums[i] - nums[j]) >= valueDifference {
                return []int{i, j}
            }
        }
    }
    return []int{-1, -1}
}