func maxOperations(nums []int) int {
    s := nums[0] + nums[1]
    result := 1
    for i := 2; i < len(nums) - 1; i += 2 {
        if nums[i] + nums[i + 1] == s {
            result += 1
        } else {
            break
        }
    }
    return result
}