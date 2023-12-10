func findKDistantIndices(nums []int, key int, k int) []int {
    result := []int{}

    for i := 0; i < len(nums); i ++ {
        if num := nums[i]; num == key {
            startIdx := i - k
            endIdx := i + k

            if startIdx < 0 {
                startIdx = 0
            }

            if len(result) > 0 && result[len(result)-1] >= startIdx {
                startIdx = result[len(result)-1] + 1
            }

            if endIdx > len(nums)-1 {
                endIdx = len(nums) - 1
            }

            for j := startIdx; j <= endIdx; j++ {
                result = append(result, j)
            }
        }
    }
    return result
}