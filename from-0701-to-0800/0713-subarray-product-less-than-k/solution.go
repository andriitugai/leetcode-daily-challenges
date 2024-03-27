func numSubarrayProductLessThanK(nums []int, k int) int {
    if k < 0 {
        return 0
    }
    result := 0
    product := 1
    left, right := 0, 0

    for right < len(nums) {
        product *= nums[right]

        for product >= k && left <= right {
            product /= nums[left]
            left += 1
        }

        result += (right - left + 1)
        right += 1
    }
    return result
}