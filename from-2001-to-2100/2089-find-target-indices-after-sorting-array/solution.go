func targetIndices(nums []int, target int) []int {
    sort.Ints(nums)
    result := []int{}
    for i := 0; i < len(nums); i ++ {
        if nums[i] == target {
            result = append(result, i)
        } else if nums[i] > target {
            break
        }
    }
    return result
}