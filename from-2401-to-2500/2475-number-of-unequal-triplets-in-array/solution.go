func unequalTriplets(nums []int) int {
    result := 0
    for i := 0; i < len(nums); i ++ {
        for j := i + 1; j < len(nums); j ++ {
            for k := j + 1; k < len(nums); k ++ {
                if nums[i] != nums[k] && nums[i] != nums[j] && nums[j] != nums[k] {
                    result += 1
                }
            }
        }
    }
    return result
}