func missingNumber(nums []int) int {
    n := len(nums)
    s := 0
    for _, num := range nums {
        s += num
    }
    return n * (n + 1) / 2 - s
}