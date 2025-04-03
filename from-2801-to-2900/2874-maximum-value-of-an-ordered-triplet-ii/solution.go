func maximumTripletValue(nums []int) int64 {
    var result int64 = 0 
    maxi, diff := -1, 0
    n := len(nums)
    for i := 0; i < n; i ++ {
        maxi = max(maxi, nums[i])
        if i >= 2 {
            result = max64(result, int64(nums[i] * diff))
        }
        if i >= 1 {
            diff = max(diff, maxi - nums[i])
        }
    }
    return result
}

func max64(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}