func alternatingSubarray(nums []int) int {
    i, ans, j := 0, -1, 1
    for j < len(nums) {
        if nums[j] == nums[i] + ((j - i) % 2) {
            ans = max(ans, j - i + 1)
            j ++
        } else {
            i ++
            j = i + 1
        }
    }
    return ans
}