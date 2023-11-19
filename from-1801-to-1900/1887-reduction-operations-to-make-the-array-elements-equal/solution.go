func reductionOperations(nums []int) int {
    sort.Ints(nums)
    result := 0
    diff := 0
    for i := 0; i < len(nums); i ++ {
        if i > 0 && nums[i] != nums[i - 1] {
            diff += 1
        }
        result += diff
    }
    return result
}