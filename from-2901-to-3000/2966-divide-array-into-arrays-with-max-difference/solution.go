func divideArray(nums []int, k int) [][]int {
    result := [][]int{}
    sort.Ints(nums)

    for p := 0; p <= len(nums) - 3; p += 3 {
        if nums[p + 2] - nums[p] <= k {
            result = append(result, nums[p:p + 3])
        } else {
            return [][]int{}
        }
    }
    return result
}