func maximumTripletValue(nums []int) int64 {
    var maxVal int64 = 0
    n := len(nums)
    nums64 := make([]int64, n)
    for i := 0; i < n; i ++ {
        nums64[i] = int64(nums[i])
    }
    
    for i := 0; i < n - 2; i ++ {
        for j := i + 1; j < n - 1; j ++ {
            for k := j + 1; k < n; k ++ {
                val := (nums64[i] - nums64[j]) * nums64[k]
                if val > maxVal {
                    maxVal = val
                }
            }
        }
    }
    return maxVal
}