func distinctAverages(nums []int) int {
    avgs := make(map[float64]bool)
    sort.Ints(nums)
    i, j := 0, len(nums) - 1
    for i < j {
        avgs[(float64(nums[i]) + float64(nums[j])) / 2.0] = true
        i += 1
        j -= 1
    }
    return len(avgs)
}