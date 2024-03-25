func findDuplicates(nums []int) []int {
    var idx int
    result := []int{}
    for _, num := range nums {
        idx = abs(num) - 1
        if nums[idx] < 0 {
            result = append(result, idx + 1)
        } else {
            nums[idx] = -nums[idx]
        }
    }
    return result
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}