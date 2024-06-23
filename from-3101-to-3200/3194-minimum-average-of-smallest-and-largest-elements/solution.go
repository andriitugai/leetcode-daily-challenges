func minimumAverage(nums []int) float64 {
    n := len(nums)
    minAvg := math.MaxFloat64
    var avg float64
    sort.Ints(nums)
    l, r := 0, n - 1
    for l < r {
        avg = float64(nums[l] + nums[r])/2.0
        if avg < minAvg {
            minAvg = avg
        }
        l += 1
        r -= 1
    }
    return minAvg
}