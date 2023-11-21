func maximizeSum(nums []int, k int) int {
    maxN := 0
    for _, num := range nums {
        if num > maxN {
            maxN = num
        }
    }
    return k * maxN + k * (k - 1) / 2
}