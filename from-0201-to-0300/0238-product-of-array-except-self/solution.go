func productExceptSelf(nums []int) []int {
    n := len(nums)
    pre := make([]int, n)
    suf := make([]int, n)

    currPre, currSuf := 1, 1
    for i, j := 0, n - 1; i < n; i, j = i + 1, j - 1 {
        pre[i] = currPre
        suf[j] = currSuf
        currPre *= nums[i]
        currSuf *= nums[j]
    }

    result := make([]int, n)
    for i := 0; i < n; i ++ {
        result[i] = pre[i] * suf[i]
    }

    return result
}