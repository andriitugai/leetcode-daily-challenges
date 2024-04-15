func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    var helper func(used int) [][]int
    
    helper = func(used int) [][]int{
        result := [][]int{}
        if used == len(nums) {
            return append(result, make([]int, 0))
        }
        for idx := 0; idx < len(nums); idx++ {
            if nums[idx] == math.MaxInt32 || (idx > 0 && nums[idx] == nums[idx - 1]) {
                continue
            }
            num := nums[idx]
            nums[idx] = math.MaxInt32
            next := helper(used + 1)
            for _, item := range next {
                prefix := append([]int{num}, item...)
                result = append(result, prefix)
            }
            nums[idx] = num
        }
        return result
    }
    return helper(0)
}