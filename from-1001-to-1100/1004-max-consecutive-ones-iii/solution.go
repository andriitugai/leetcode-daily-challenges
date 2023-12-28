func longestOnes(nums []int, k int) int {
    left, right := 0, 0
    result := 0
    zeros := 0

    for i := 0; i < len(nums); i ++ {
        if nums[i] == 0 {
            zeros += 1
        }
        for zeros > k {
            if nums[left] == 0 {
                zeros -= 1
            }
            left += 1
        }

        right += 1
        if result < right - left {
            result = right - left
        }
    }
    return result
}