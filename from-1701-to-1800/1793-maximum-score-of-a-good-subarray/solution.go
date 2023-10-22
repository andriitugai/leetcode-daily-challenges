func maximumScore(nums []int, k int) int {
    left, right := k, k
    minVal := nums[k]
    maxScore := minVal

    for left > 0 || right < len(nums) - 1 {
        // Decide which direction to expand based on values and boundaries
        if left == 0 || (right < len(nums) - 1 && nums[right + 1] > nums[left - 1]) {
            right += 1
        } else {
            left -= 1
        }

        // Update min_val and max_score
        if nums[left] < minVal {
            minVal = nums[left]
        }
        if nums[right] < minVal {
            minVal = nums[right]
        }
        if minVal * (right - left + 1) > maxScore {
            maxScore = minVal * (right - left + 1)
        }
    }
    return maxScore
}