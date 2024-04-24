func findValueOfPartition(nums []int) int {
    sort.Ints(nums)
    result := 1000000001
    var diff int
    for pi := 0; pi < len(nums) - 1; pi ++ {
        diff = nums[pi + 1] - nums[pi]
        if diff < result {
            result = nums[pi + 1] - nums[pi]
        }
    }
    return result
}