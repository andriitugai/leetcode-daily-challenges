func minPairSum(nums []int) int {
    sort.Ints(nums)
    result := 0
    i, j := 0, len(nums) - 1
    for i < j {
        s := nums[i] + nums[j]
        if s > result {
            result = s
        }
        i += 1
        j -= 1
    }
    return result
}