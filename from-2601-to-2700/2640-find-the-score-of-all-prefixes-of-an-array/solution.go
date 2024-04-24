func findPrefixScore(nums []int) []int64 {
    var score int64
    result := make([]int64, len(nums))
    curMax := 0
    for i := 0; i < len(nums); i ++ {
        if nums[i] > curMax {
            curMax = nums[i]
        }
        score += int64(nums[i] + curMax)
        result[i] = score
    }
    return result
}