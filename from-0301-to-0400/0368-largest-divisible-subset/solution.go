func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    n := len(nums)
    result := make([][]int, n)
    for i := 0; i < n; i ++ {
        result[i] = []int{nums[i]}
    }

    for i := 0; i < n; i ++ {
        for j := 0; j < i; j ++ {
            if nums[i] % nums[j] == 0 && len(result[i]) < len(result[j]) + 1 {
                copy(result[i], result[j])
                result[i] = append(result[i], nums[i])
            }
        }
    }
    maxLen := -1
    ans := []int{}
    for i := 0; i < n; i ++ {
        if len(result[i]) > maxLen {
            maxLen = len(result[i])
            ans = result[i]
        }
    }
    return ans
}